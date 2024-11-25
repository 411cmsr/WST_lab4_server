package services

import (
	"WST_lab4_server/internal/database"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router   *mux.Router
	logger   *logrus.Logger
	database database.Database
}

func newServer(database database.Database) *server {
	s := &server{
		router:   mux.NewRouter(),
		logger:   logrus.New(),
		database: database,
	}
	s.configureRouter()
	return s
}
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)

}

func (s *server) configureRouter() {
	///
}
