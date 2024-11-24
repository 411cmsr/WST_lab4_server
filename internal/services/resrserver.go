package services

import (
	"WST_lab4_server/internal/database"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type RESTServer struct {
	config   *Config
	logger   *logrus.Logger
	router   *mux.Router
	database *database.Database
}

func New(config *Config) *RESTServer {
	return &RESTServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

}
func (s *RESTServer) Run() error {
	if err := s.configLogger(); err != nil {
		return err
	}
	s.configRouter()
	if err := s.configDatabase(); err != nil {
		return err
	}
	s.logger.Info("REST server started")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *RESTServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *RESTServer) configRouter() {
	s.router.HandleFunc("/hello", s.handleHello())

}
func (s *RESTServer) configDatabase() error {
	st := database.New(s.config.DataBase)
	if err := st.Open(); err != nil {
		return err
	}
	s.database = st
	return nil
}

func (s *RESTServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!!!!!")
	}
}
