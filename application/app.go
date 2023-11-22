package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRouter(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    ":8000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Errorf("Server Start failed: %w", err)
	}
	return nil
}
