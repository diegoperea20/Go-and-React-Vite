package main

import (
	"gorm/handlers"
	"gorm/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	models.MigrarTask()

	//Rutas
	mux := mux.NewRouter()

	// Middleware CORS con opciones personalizadas
    corsOptions := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:5173"}, // Reemplaza con la URL de tu frontend React
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    })
    handler := corsOptions.Handler(mux)

    // Endpoints
    mux.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
    mux.HandleFunc("/tasks/{id:[0-9]+}", handlers.GetTask).Methods("GET")
    mux.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
    mux.HandleFunc("/tasks/{id:[0-9]+}", handlers.UpdateTask).Methods("PUT")
    mux.HandleFunc("/tasks/{id:[0-9]+}", handlers.DeleteTask).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":5000", handler))

}


/* Comands for use docker container mysql
docker run --name mymysql -e MYSQL_ROOT_PASSWORD=mypassword -p 3306:3306 -d mysql:latest
docker exec -it mymysql bash
mysql -u root -p
create database gomysql; */