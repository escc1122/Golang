package observer

import (
	"context"
	"dig_test/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"time"
)

type IServer interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type digIn struct {
	dig.In
	Log utils.Log
}

type GinServer struct {
	srv *http.Server
	in  digIn
}

func (g *GinServer) Run(ctx context.Context) error {
	e := gin.New()

	//測試關機是否會等待
	e.GET("test", func(context *gin.Context) {
		<-time.After(25 * time.Second)
		context.Data(http.StatusOK, "application/json; charset=utf-8", []byte("ok"))
	})

	g.srv = &http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		g.in.Log.Info("listen: %s\n", err)
		return err
	}

	return nil
}

func (g *GinServer) Shutdown(ctx context.Context) error {
	if err := g.srv.Shutdown(ctx); err != nil {
		return g.in.Log.Error("Server Shutdown: %v\n", err)
	}

	return nil
}

type GinServer2 struct {
	srv *http.Server
	in  digIn
}

func (g *GinServer2) Run(ctx context.Context) error {
	e := gin.New()
	e.GET("test", func(context *gin.Context) {
		<-time.After(5 * time.Second)
		context.Data(http.StatusOK, "application/json; charset=utf-8", []byte("ok"))
	})

	g.srv = &http.Server{
		Addr:    ":8081",
		Handler: e,
	}

	if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		g.in.Log.Info("listen: %s\n", err)
		return err
	}

	return nil
}

func (g *GinServer2) Shutdown(ctx context.Context) error {
	if err := g.srv.Shutdown(ctx); err != nil {
		return g.in.Log.Error("Server2 Shutdown: %v\n", err)
	}

	return nil
}
