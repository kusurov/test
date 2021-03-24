package src

import (
	"awesomeProject2/pkg/internal"
	"context"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func Start(configPath string, loggingPath string) error {
	config, err := internal.NewConfig(configPath)
	if err != nil {
		return err
	}

	sessionStore := sessions.NewCookieStore([]byte(config.Api.SessionKey))

	srv := internal.NewServer(config, sessionStore)
	if err := srv.InitializeLogging(loggingPath); err != nil {
		return err
	}

	if err := srv.CreateStore(); err != nil {
		return err
	}

	routerHandler(srv)

	server := &http.Server{
		Addr:    ":" + config.Api.BindAddr,
		Handler: srv,
	}

	go func() {
		srv.Logger.Info("Starting API server on port: " + config.Api.BindAddr)

		if err := server.ListenAndServe(); err != nil {
			srv.Logger.Error(err)
		}
	}()

	cancel := make(chan os.Signal, 1)
	signal.Notify(cancel, os.Interrupt)

	<-cancel
	srv.Logger.Info("Closing server: " + config.Api.BindAddr)
	if err := server.Shutdown(context.Background()); err != nil {
		return err
	}

	if err := srv.CloseDB(); err != nil {
		return err
	}

	log.Println("Server stopped!")

	return nil
}
