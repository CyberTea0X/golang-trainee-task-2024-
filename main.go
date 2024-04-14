package main

import (
	"gobanners/controllers"
	"gobanners/models"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbconf, err := models.DBConfigFromEnv(".env")
	if err != nil {
		panic(err)
	}
	log.Println("Connecting to the database..")
	db, err := models.SetupDatabase(dbconf)
	if err != nil {
		panic(err)
	}
	for err := db.Ping(); err != nil; err = db.Ping() {
		log.Println(err)
		time.Sleep(time.Second)
		log.Println("Database ping failed, retrying..")
	}
	log.Println("Successfully connected to the database")
	err = models.MigrateDatabase(db)
	if err != nil {
		panic(err)
	}
	pCtrl := controllers.NewPublicController(db)
	gin := controllers.SetupRouter(pCtrl)
	log.Println("Gobanners starting")
	err = gin.Run()
	if err != nil {
		panic(err)
	}
}
