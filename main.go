package main

import (
	"golang-rest-user/database"
	"golang-rest-user/handler"

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
	masterDB := database.ConnectMasterDB()
	if err := database.InitTenantDBs(masterDB); err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	tntRepo := repository.NewTenantRepo(masterDB)
	tntSvc := service.NewTenantService(tntRepo)
	tntHandler := handler.NewTenantHandler(tntSvc)

	userHandler := handler.NewUserHandler()

	authHandler := handler.NewAuthHandler()

	routes.RegisterRoutes(r, userHandler, tntHandler, authHandler)

	r.Run(":8080")

}
