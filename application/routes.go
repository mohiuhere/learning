package application

import (
	"test/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadUserRoutes(router chi.Router) {
	userHandler := &handler.User{}

	//sub-router for User
	router.Post("/", userHandler.CreateUser)
	router.Get("/", userHandler.ListUser)
	router.Get("/{id}", userHandler.GetUserByID)
	router.Put("/{id}", userHandler.UpdateUserById)
	router.Get("/{id}", userHandler.DeleteUserById)
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	//sub-routes for Order
	router.Post("/", orderHandler.CreateOrder)
	router.Get("/", orderHandler.ListOrder)
	router.Get("/{id}", orderHandler.GetOrderById)
	router.Put("/{id}", orderHandler.UpdateOrderById)
	router.Delete("/{id}", orderHandler.DeleteOrderById)
}

func loadRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	//Root routes for sub-route
	router.Route("/order", loadOrderRoutes)
	router.Route("/user", loadUserRoutes)

	return router
}
