package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
	"log"
	"net/http"
)

type Admin struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Heslo string `json:"heslo"`
}

// Create User Table
func CreateAdminTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Admin{}, opts)
	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Runner table created")
	return nil
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllAdmins(c *gin.Context) {
	var todos []Admin
	err := dbConnect.Model(&todos).Select()

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    todos,
	})
	return
}

func CreateAdmin(c *gin.Context) {
	var todo Admin
	c.BindJSON(&todo)
	login := todo.Login
	heslo := todo.Heslo
	id := guuid.New().String()

	insertError := dbConnect.Insert(&Admin{
		ID:    id,
		Login: login,
		Heslo: heslo,
	})
	if insertError != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Runner created Successfully",
	})
	return
}

func GetSingleAdmin(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Admin{ID: todoId}
	err := dbConnect.Select(todo)

	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Runner not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Runner",
		"data":    todo,
	})
	return
}

func EditAdmin(c *gin.Context) {
	todoId := c.Param("todoId")
	meno := c.Param("meno")
	var todo Admin
	c.BindJSON(&todo)

	_, err := dbConnect.Model(&Admin{}).Set("meno = ?", meno).Where("id = ?", todoId).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Runner Edited Successfully",
	})
	return
}

func DeleteAdmin(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Admin{ID: todoId}

	err := dbConnect.Delete(todo)
	if err != nil {
		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Runner deleted successfully",
	})
	return
}
