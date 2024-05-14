package main

import (
	"net/http"

	"github.com/sathish-30/book-management/internal/config"
	db "github.com/sathish-30/book-management/internal/database"
	"github.com/sathish-30/book-management/internal/handler"
	"github.com/sathish-30/book-management/internal/middleware"
)

func init() {
	config.Configuration.LoadConfig()
	config.Configuration.DisplayConfig()
	db.ConnectToDatabase()
}

func main() {
	PORT := ":" + config.Configuration.PORT
	app := http.NewServeMux()
	app.HandleFunc("GET /v1/healthcheck", handler.HealthcheckHandler)
	app.Handle("POST /v1/signup", middleware.SetContentTypeJson(http.HandlerFunc(handler.SignUpHandler)))
	http.ListenAndServe(PORT, app)
}
