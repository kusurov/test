package app

import (
	"context"
	"github.com/gorilla/sessions"
	"kusurovAPI/internal/configs"
	"kusurovAPI/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func Start(configPath string, loggingPath string) error {
	config, err := configs.NewConfig(configPath)
	if err != nil {
		return err
	}

	sessionStore := sessions.NewCookieStore([]byte(config.Api.SessionKey))

	srv := server.NewServer(config, sessionStore)
	if err := srv.InitializeLogging(loggingPath); err != nil {
		return err
	}

	if err := srv.CreateStore(); err != nil {
		return err
	}

	routerHandler(srv)

	initServer := &http.Server{
		Addr:    ":" + config.Api.BindAddr,
		Handler: srv,
	}

	go func() {
		srv.Logger.Info("Starting API server on port: " + config.Api.BindAddr)

		if err := initServer.ListenAndServe(); err != nil {
			srv.Logger.Error(err)
		}
	}()

	cancel := make(chan os.Signal, 1)
	signal.Notify(cancel, os.Interrupt)

	<-cancel
	srv.Logger.Info("Closing server: " + config.Api.BindAddr)
	if err := initServer.Shutdown(context.Background()); err != nil {
		return err
	}

	if err := srv.CloseDB(); err != nil {
		return err
	}

	log.Println("Server stopped!")

	return nil
}
