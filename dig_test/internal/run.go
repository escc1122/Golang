package internal

import (
	"context"
	"dig_test/internal/binder"
	"dig_test/internal/server/observer"
	"dig_test/internal/server/subject"
	"fmt"
	"os"
)

func Run() {
	binder := binder.GetDigInstance()
	if err := binder.Invoke(startServer); err != nil {
		panic(err)
	}
}

func startServer() {
	ctx := context.Background()
	si := new(subject.ServerInterrupt)
	si.Listen()

	servers := []observer.IServer{
		new(observer.GinServer),
		new(observer.GinServer2),
	}

	for _, server := range servers {
		si.Register(server)
		go func(server observer.IServer) {
			if err := server.Run(ctx); err != nil {
				fmt.Printf("run error: %v \n", err)
				os.Exit(1)
			}
		}(server)
	}

	select {}
}
