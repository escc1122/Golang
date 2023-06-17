package subject

import (
	"context"
	"dig_test/internal/server/observer"
	"dig_test/internal/utils"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"sync"
	"time"
)

type digIn struct {
	dig.In
	Log utils.Log
}

type ServerInterrupt struct {
	mx      sync.Mutex
	servers []observer.IServer
	in      digIn
}

func (s *ServerInterrupt) Register(server observer.IServer) {
	s.mx.Lock()
	s.servers = append(s.servers, server)
	s.mx.Unlock()
}

func (s *ServerInterrupt) Listen() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go func() {
		c := <-interrupt

		s.mx.Lock()
		defer s.mx.Unlock()
		s.in.Log.Info("Server Shutdown, osSignal: %v\n", c)
		s.in.Log.Info("Shutdown Server ...\n")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		errs := make(chan error, len(s.servers))
		var wg sync.WaitGroup
		wg.Add(len(s.servers))
		for _, server := range s.servers {
			go func(server observer.IServer) {
				defer wg.Done()
				errs <- server.Shutdown(ctx)
			}(server)
		}

		wg.Wait()
		close(errs)

		var isException bool

		for err := range errs {
			if err != nil {
				isException = true
				s.in.Log.Info("Shutdown error:", err)
			}
		}

		s.in.Log.Info("Server exiting\n")
		if isException {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}()
}
