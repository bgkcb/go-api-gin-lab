package main

import (
	"github.com/gin-gonic/gin"

	"go-api-gin-lab/config"
	"go-api-gin-lab/handlers"
	"go-api-gin-lab/repositories"
	"go-api-gin-lab/services"
)

func main() {
	db := config.InitDB()

	repo := &repositories.StudentRepository{DB: db}
	service := &services.StudentService{Repo: repo}
	handler := &handlers.StudentHandler{Service: service}

	r := gin.Default()

	r.GET("/students", handler.GetStudents)
	r.GET("/students/:id", handler.GetStudentByID)
	r.POST("/students", handler.CreateStudent)

	r.Run(":8080")
}
