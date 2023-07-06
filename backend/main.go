package main

import (
	"log"
	"net/http"
	"strconv"

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

var db *gorm.DB

func main() {
	//connect to db
	var err error
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
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

//getALL
func getTasks(c *gin.Context) {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

//getOne
func getTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error" : err.Error()})
		return
	}

	var task Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var task Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		return
	}

	var updatedTask Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	task.Title = updatedTask.Title
	task.Completed = updatedTask.Completed

	if err := db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var task Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		return
	}

	if err := db.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message" : "Task deleted successfully"})
}