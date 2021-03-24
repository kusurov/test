package internal

import (
	"awesomeProject2/pkg/internal/store"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	Config       *Config
	Router       *mux.Router
	Logger       *logrus.Logger
	Store        *store.Store
	SessionStore sessions.Store

	db *sql.DB
}

func NewServer(config *Config, sessionStore sessions.Store) *Server {
	s := &Server{
		Config:       config,
		Router:       mux.NewRouter(),
		Logger:       logrus.New(),
		SessionStore: sessionStore,
	}

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) CreateStore() error {
	db, err := sql.Open("mysql", s.Config.Database.Username+":"+s.Config.Database.Password+"@tcp("+s.Config.Database.Host+")/"+s.Config.Database.Dbname)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.Logger.Info("Connected to MySQL table")

	s.db = db
	s.Store = store.New(db, s.Logger)

	return nil
}

func (s *Server) CloseDB() error {
	s.Logger.Info("Disconnected database")

	return s.db.Close()
}

func (s *Server) InitializeLogging(loggingPath string) error {
	loggingName := "logs" + time.Now().Format("20060102150405030405") + ".log"

	if err := os.MkdirAll(loggingPath, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(filepath.Join(loggingPath, loggingName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	s.Logger.SetOutput(io.MultiWriter(os.Stdout, file))
	s.Logger.SetFormatter(&logrus.JSONFormatter{})

	return nil
}
