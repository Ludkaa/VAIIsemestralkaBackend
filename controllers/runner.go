package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
	"log"
	"net/http"
)

type Runner struct {
	ID         string `json:"id"`
	Meno       string `json:"meno"`
	Priezvisko string `json:"priezvisko"`
	Birthday   string `json:"birthday"`
	Email      string `json:"email"`
	Trat       string `json:"trat"`
	Suhlas     string `json:"suhlas"`
}

// Create User Table
func CreateTodoTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Runner{}, opts)
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

func GetAllTodos(c *gin.Context) {
	var todos []Runner
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

func CreateTodo(c *gin.Context) {
	var todo Runner
	c.BindJSON(&todo)
	meno := todo.Meno
	priezvisko := todo.Priezvisko
	birthday := todo.Birthday
	email := todo.Email
	trat := todo.Trat
	suhlas := todo.Suhlas
	id := guuid.New().String()

	insertError := dbConnect.Insert(&Runner{
		ID:         id,
		Meno:       meno,
		Priezvisko: priezvisko,
		Birthday:   birthday,
		Email:      email,
		Trat:       trat,
		Suhlas:     suhlas,
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

func GetSingleTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Runner{ID: todoId}
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

func EditTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	meno := c.Param("meno")
	var todo Runner
	c.BindJSON(&todo)

	_, err := dbConnect.Model(&Runner{}).Set("meno = ?", meno).Where("id = ?", todoId).Update()
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

func DeleteTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Runner{ID: todoId}

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
