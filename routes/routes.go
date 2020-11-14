package routes

import (
	"backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/runners", controllers.GetAllTodos)
	router.POST("/runner", controllers.CreateTodo)
	router.GET("/runner/:todoId", controllers.GetSingleTodo)
	router.PUT("/runner/:todoId", controllers.EditTodo)
	router.DELETE("/runner/:todoId", controllers.DeleteTodo)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
