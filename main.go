package main

import (
	"gobanner/controllers"
	"gobanner/models"
)

func main() {
	dbconf, err := models.DBConfigFromEnv(".env")
	if err != nil {
		panic(err)
	}
	db, err := models.SetupDatabase(dbconf)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	pCtrl := controllers.NewPublicController(db)
	gin := controllers.SetupRouter(pCtrl)
	gin.Run()
}
