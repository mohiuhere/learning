package application

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
	db     *sql.DB
}

func New() *App {

	db, err := sql.Open("postgres", "user=root password=123456 dbname=test sslmode=disable")
	if err != nil {
		// Handle error, e.g., log it and return or panic
		fmt.Println("Error opening database connection:", err)
	}

	app := &App{
		router: loadRouter(),
		rdb:    redis.NewClient(&redis.Options{}),
		db: db,
	}
	return app
}

func (a *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    ":8000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to conect redis: %w", err)
	}

	err = a.db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect db: %w", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}
