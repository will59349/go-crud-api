package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/will59349/go-crud-api/handler"
	"github.com/will59349/go-crud-api/pkg/database"
	"log"
	"os"

	_ "github.com/jmoiron/sqlx"
)

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can't find .env file")
	}

	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbname)

	database.DB, err = database.ConnectDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to MySQL!")
}

func main() {
	// Initialize database
	InitDB()

	// Initialize Gin router
	r := gin.Default()

	// Define routes
	r.GET("/users", handler.GetUsersHandler)
	r.GET("/users/:id", handler.GetUserHandler)
	r.POST("/users", handler.CreateUserHandler)
	r.PUT("/users/:id", handler.UpdateUserHandler)
	r.DELETE("/users/:id", handler.DeleteUserHandler)
	r.PATCH("/users/:id", handler.DynamicUpdateUserHandler)

	// Start the server
	r.Run(":8080")
}
