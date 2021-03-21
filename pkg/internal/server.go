package internal

import (
	"awesomeProject2/pkg/internal/store"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	Config  	*Config
	Router		*mux.Router
	Logger		*logrus.Logger
	Store		*store.Store
	SessionStore sessions.Store
}

func NewServer(config *Config, sessionStore sessions.Store) *Server {
	s := &Server{
		Config: config,
		Router: mux.NewRouter(),
		Logger: logrus.New(),
		SessionStore: sessionStore,
	}

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) CreateStore() error {
	db, err := sql.Open("mysql", s.Config.Database.Username + ":" + s.Config.Database.Password + "@tcp(" + s.Config.Database.Host + ")/" + s.Config.Database.Dbname)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.Logger.Info("Connected to MySQL table")

	s.Store = store.New(db, s.Logger)

	return nil
}