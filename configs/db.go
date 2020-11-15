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

	urlOs := os.Getenv("DATABASE_URL")
	if urlOs == "" {
		pgOptions, err := pg.ParseURL("postgresql://dyeeifmazjyqcn:76dcd0cd67541915e489d680de407a089482b0f298684bab895f1ab08085aafd@ec2-54-246-115-40.eu-west-1.compute.amazonaws.com:5432/d7aar8r8v5qv2a?sslmode=require")
		if err != nil {
			panic(err)
		}
		var db = pg.Connect(pgOptions)
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
	} else {
		parsedUrl, err := url.Parse(urlOs)
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
		var db = pg.Connect(pgOptions)
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
}
