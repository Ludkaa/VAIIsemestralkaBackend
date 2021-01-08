package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
	"log"
	"net/http"
)

type Referencie struct {
	ID    string `json:"id"`
	Meno  string `json:"meno"`
	Email string `json:"email"`
	Text  string `json:"text"`
}

// Create User Table
func CreateReferencieTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Referencie{}, opts)
	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Referencie table created")
	return nil
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbRefConnect *pg.DB

func InitiateRefDB(db *pg.DB) {
	dbRefConnect = db
}

func GetAllReferencie(c *gin.Context) {
	var todos []Referencie
	err := dbRefConnect.Model(&todos).Select()

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

func CreateReferencie(c *gin.Context) {
	var todo Referencie
	c.BindJSON(&todo)
	meno := todo.Meno
	email := todo.Email
	text := todo.Text
	id := guuid.New().String()

	insertError := dbRefConnect.Insert(&Referencie{
		ID:    id,
		Meno:  meno,
		Email: email,
		Text:  text,
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

func GetSingleReferencie(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Referencie{ID: todoId}
	err := dbRefConnect.Select(todo)

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

func EditReferencie(c *gin.Context) {
	var todo Referencie
	c.ShouldBindJSON(&todo)
	meno := todo.Meno
	email := todo.Email
	text := todo.Text
	todoId := c.Param("todoId")

	_, err := dbRefConnect.Model(&Referencie{}).Set("meno = ?,email = ?, text = ?", meno, email, text).Where("id = ?", todoId).Update()
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

func DeleteReferencie(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Referencie{ID: todoId}

	err := dbRefConnect.Delete(todo)
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
		"message": "Referencie deleted successfully",
	})
	return
}
