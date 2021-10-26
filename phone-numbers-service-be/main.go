package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"phone-numbers-service.jpay.task/src/modules/countries"
	"phone-numbers-service.jpay.task/src/modules/customers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	db, err := gorm.Open(sqlite.Open("./db/sample.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	countryRepo := countries.NewCountryRepository()
	customerRepo := customers.NewCustomerRepository(db)

	countryRouter := countries.NewCountryRouter(countryRepo)
	customerRouter := customers.NewCustomerRouter(customerRepo, countryRepo)

	router := gin.Default()

	router.GET("/countries", countryRouter.ListCountries)
	router.GET("/customers", customerRouter.ListCustomers)

	router.Run(fmt.Sprintf(":%s", port))
}
