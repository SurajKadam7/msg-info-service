package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/log"
	"github.com/jackc/pgx/v5/pgxpool"
	msginfosrv "github.com/suraj.kadam7/msg-info-srv/msginfo_srv"
	"github.com/suraj.kadam7/msg-info-srv/msginfo_srv/service"
	"github.com/suraj.kadam7/msg-info-srv/msginfo_srv/transport"
	transporthttp "github.com/suraj.kadam7/msg-info-srv/msginfo_srv/transport/http"
	"github.com/suraj.kadam7/msg-info-srv/repos/msginfo/postgres"
)

func main() {
	var client *pgxpool.Pool
	{
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		// urlExample := "postgres://username:password@localhost:5432/database_name"
		client, err = pgxpool.New(ctx, "")
		if err != nil {
			panic("Todo")
		}
		defer client.Close()
	}

	logger := log.With(log.NewJSONLogger(os.Stdout), "service", "msg-info-srv")

	repo := postgres.New(client, "")
	srv := msginfosrv.New(repo)
	srv = service.LoggingMiddleware(logger)(srv)
	endpoints := transport.Endpoints(srv)
	handler := transporthttp.NewHTTPHandler(&endpoints)

	// server code
	httpServer := http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	// gracefull shutdown ..
	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := httpServer.ListenAndServe()
		// error will never be nil here ...
		logger.Log("server error ", err.Error())
		close(quite)
	}()

	sig := <-quite

	if sig == syscall.SIGINT || sig == syscall.SIGTERM {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		logger.Log("server status", "shutdown called")
		time.Sleep(time.Second * 5)

		err := httpServer.Shutdown(ctx)
		if err != nil {
			logger.Log("error while shutdown", err)
		}
	}

	logger.Log("server status", "closed")

}
