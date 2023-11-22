package handler

import (
	"fmt"
	"net/http"
)

type Order struct {}

func (o *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Order")
}

func (o *Order) ListOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Read Order")
}

func (o *Order) GetOrderById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Read Order By Id")
}

func (o *Order) UpdateOrderById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Order By Id")
}

func (o *Order) DeleteOrderById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Order By Id")
}
