package routes

import (
	"backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/admins", controllers.GetAllAdmins)
	router.POST("/admin", controllers.CreateAdmin)
	router.POST("/adminlogin", controllers.AdminLogin)
	router.GET("/admin/:todoId", controllers.GetSingleAdmin)
	router.PUT("/admin/:todoId", controllers.EditAdmin)
	router.DELETE("/admin/:todoId", controllers.DeleteAdmin)

	router.GET("/runners", controllers.GetAllRunners)
	router.POST("/runner", controllers.CreateRunner)
	router.GET("/runner/:todoId", controllers.GetSingleRunner)
	router.PUT("/runner/:todoId", controllers.EditRunner)
	router.DELETE("/runner/:todoId", controllers.DeleteRunner)

	router.GET("/referencie", controllers.GetAllReferencie)
	router.POST("/referencia", controllers.CreateReferencie)
	router.GET("/referencia/:todoId", controllers.GetSingleReferencie)
	router.PUT("/referencia/:todoId", controllers.EditReferencie)
	router.DELETE("/referencia/:todoId", controllers.DeleteReferencie)
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
