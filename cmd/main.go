package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	go_crud_api "github.com/will59349/go-crud-api/api"
	"log"
	"net/http"
)

func main() {
	go_crud_api.InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/users", go_crud_api.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", go_crud_api.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", go_crud_api.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", go_crud_api.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", go_crud_api.DynamicUpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", go_crud_api.DeleteUserHandler).Methods("DELETE")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
