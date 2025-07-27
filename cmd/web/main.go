package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/scrivx/go-api-web/cmd/web/routes"
	"github.com/scrivx/go-api-web/cmd/web/routes/mainRoute"
)

func InitRoutes() {
	mainRoute.MainRender()
	mainRoute.ClockRender()
}

func main() {
	InitRoutes()
	mux := routes.GetMuxInstance()
	fileServer := routes.GetFileServerInstance()

	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	slog.Log(context.TODO(), slog.LevelInfo, "🚀 Starting server on http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	
}