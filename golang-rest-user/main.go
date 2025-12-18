package main

import (
	"golang-rest-user/database"
	"golang-rest-user/handler"

	//"golang-rest-user/middleware"

	//"golang-rest-user/models"
	"golang-rest-user/repository"
	"golang-rest-user/routes"
	"golang-rest-user/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.ConnectMasterDB()
	DB := database.DB
	if err := database.InitTenantDBs(database.DB); err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	tntRepo := repository.NewTenantRepo(DB)
	tntSvc := service.NewTenantService(tntRepo)
	tntHandler := handler.NewTenantHandler(tntSvc)

	userHandler := handler.NewUserHandler()

	routes.RegisterRoutes(r, userHandler, tntHandler)

	r.Run(":8080")

}
