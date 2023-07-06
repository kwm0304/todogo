package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Task struct {
	gorm.Model
	Title			string
	Completed	bool
	UserId		uint
}

func main() {
	//connect to db
	db, err := gorm.Open("sqlite3", "todo.db")
	if err != nil {
		log.Fatal("Unable to connect to db:", err)
	}
	defer db.Close()

	//auto-migrate the trask struct to create the 'task' table
	db.AutoMigrate(&Task{})

	//create new gin router
	router := gin.Default()

	//enable cors middleware
	router.Use(cors.Default())

	//define routes
	router.POST("/tasks", createTask)
	router.GET("/tasks", getTasks)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	//starting http server
	log.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//gin.Context encapsulates the req and res for requests
func createTask(c *gin.Context) {

}

func getTasks(c *gin.Context) {

}

func updateTask(c *gin.Context) {

}

func deleteTask(c *gin.Context) {

}