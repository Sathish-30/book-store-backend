package main

import (
	"encoding/json"
	"net/http"

	"github.com/sathish-30/book-management/internal/config"
	db "github.com/sathish-30/book-management/internal/database"
)

func init() {
	config.Configuration.LoadConfig()
	config.Configuration.DisplayConfig()
	db.ConnectToDatabase()
}

func main() {
	PORT := ":" + config.Configuration.PORT
	app := http.NewServeMux()
	app.HandleFunc("GET /v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{
			"Message": "Health check",
		})
	})
	http.ListenAndServe(PORT, app)
}
