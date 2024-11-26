package services

import (
	"WST_lab4_server/internal/database/sqldb"
	"database/sql"
	"github.com/gorilla/sessions"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	database := sqldb.New(db)
	sessionsStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(database, sessionsStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
