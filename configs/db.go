package config

import (
	"backend/controllers"
	"github.com/go-pg/pg/v9"
	"log"
	"net/url"
	"os"
)

// Connecting to db
func Connect() *pg.DB {

	parsedUrl, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	pgOptions := &pg.Options{
		User:     parsedUrl.User.Username(),
		Database: parsedUrl.Path[1:],
		Addr:     parsedUrl.Host,
	}

	if password, ok := parsedUrl.User.Password(); ok {
		pgOptions.Password = password
	}

	var db *pg.DB = pg.Connect(pgOptions)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateRunnerTable(db)
	controllers.InitiateRDB(db)
	controllers.CreateAdminTable(db)
	controllers.InitiateDB(db)
	return db
}
