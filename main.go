package main

import (
	"log"
	"net/http"
	"project-app-inventory-restapi-golang-faisal/config"
	"project-app-inventory-restapi-golang-faisal/database"
	"project-app-inventory-restapi-golang-faisal/routes"
	"project-app-inventory-restapi-golang-faisal/utils"

	"go.uber.org/zap"
)


func main()  {
	db := database.InitDb()
	defer func ()  {
		if err := db.Close(); err != nil {
			log.Printf("Error closing Database: %v", err)
		}
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	logger, err := utils.InitLogger()
	if err != nil {
		log.Fatal("failed to init logger: ", err)
	}
	defer logger.Sync()

	utils.InitValidator()

	r := routes.SetUpRouter()

	port := cfg.AppPort
	logger.Info("Starting server", zap.String("port", port))
	err = http.ListenAndServe("0.0.0.0:" +port, r)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
	
}