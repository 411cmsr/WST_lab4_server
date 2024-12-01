package main

import (
	"WST_lab4_server/internal/database/postgres"
	"WST_lab4_server/internal/services"
	"go.uber.org/zap"
	"net/http"
)

type MyHandler struct{}

func main() {
	services.InitializeLogger()

	configFile := "config/config.yaml"

	err := postgres.InitDB(configFile)
	if err != nil {
		services.Logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	services.Logger.Info("Database connection established successfully.")

	err = postgres.UpdateDB(configFile)
	if err != nil {
		services.Logger.Fatal("Failed to update database", zap.Error(err))
	}
	services.Logger.Info("Database updated successfully.")

	err = http.ListenAndServe(":8088", MyHandler{})
	if err != nil {
		services.Logger.Fatal("Failed to start server", zap.Error(err))
	}

}
func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
