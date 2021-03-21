package src

import (
	"awesomeProject2/pkg/internal"
	"github.com/gorilla/sessions"
	"net/http"
)

func Start(configPath string) error {
	config, err := internal.NewConfig(configPath)
	if err != nil {
		return err
	}

	sessionStore := sessions.NewCookieStore([]byte(config.Api.SessionKey))

	srv := internal.NewServer(config, sessionStore)
	if err := srv.CreateStore(); err != nil {
		return err
	}

	routerHandler(srv)

	srv.Logger.Info("Starting API server on port: " + config.Api.BindAddr)
	return http.ListenAndServe(":" + config.Api.BindAddr, srv)
}
