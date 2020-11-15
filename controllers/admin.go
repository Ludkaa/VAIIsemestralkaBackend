package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	log.Printf("Admin table created")
	return nil
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllAdmins(c *gin.Context) {
	var admins []Admin
	err := dbConnect.Model(&admins).Select()

	if err != nil {
		log.Printf("Error while getting all admins, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    admins,
	})
	return
}

func CreateAdmin(c *gin.Context) {
	var admin Admin
	c.ShouldBindJSON(&admin)
	login := admin.Login
	heslo := admin.Heslo
	id := guuid.New().String()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(heslo), 8)

	insertError := dbConnect.Insert(&Admin{
		ID:    id,
		Login: login,
		Heslo: string(hashedPassword),
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
		"message": "Admin created Successfully",
	})
	return
}

func AdminLogin(c *gin.Context) {
	var admin Admin
	c.ShouldBindJSON(&admin)
	login := admin.Login
	heslo := admin.Heslo

	admin2 := new(Admin)
	err := dbConnect.Model(admin2).Where("login = ?", login).Select()

	if err = bcrypt.CompareHashAndPassword([]byte(admin2.Heslo), []byte(heslo)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "bad",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "ok",
	})

}

func GetSingleAdmin(c *gin.Context) {
	todoId := c.Param("todoId")
	admin := &Admin{ID: todoId}
	err := dbConnect.Select(admin)

	if err != nil {
		log.Printf("Error while getting a single admin, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Admin not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Admin",
		"data":    admin,
	})
	return
}

func EditAdmin(c *gin.Context) {
	adminId := c.Param("adminId")
	meno := c.Param("meno")
	var admin Admin
	c.ShouldBindJSON(&admin)

	_, err := dbConnect.Model(&Admin{}).Set("meno = ?", meno).Where("id = ?", adminId).Update()
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
		"message": "Admin Edited Successfully",
	})
	return
}

func DeleteAdmin(c *gin.Context) {
	adminId := c.Param("adminId")
	admin := &Admin{ID: adminId}

	err := dbConnect.Delete(admin)
	if err != nil {
		log.Printf("Error while deleting a single admin, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Admin deleted successfully",
	})
	return
}
