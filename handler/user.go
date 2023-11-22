package handler

import (
	"fmt"
	"net/http"
)

type User struct{}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create User")
}

func (u *User) ListUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Read User")
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Read User By Id")
}

func (u *User) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update User By Id")
}

func (u *User) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete User By Id")
}
