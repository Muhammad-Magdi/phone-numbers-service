package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"phone-numbers-service.jpay.task/src/modules/countries"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	countryRepo := countries.NewCountryRepository()

	countryRouter := countries.NewCountryRouter(countryRepo)

	router := gin.Default()

	router.GET("/countries", countryRouter.ListCountries)

	router.Run(fmt.Sprintf(":%s", port))
}
