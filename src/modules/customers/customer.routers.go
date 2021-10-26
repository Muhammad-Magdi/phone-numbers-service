package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"phone-numbers-service.jpay.task/src/modules/countries"
)

type CustomerRouter struct {
	service CustomerServiceI
}

func NewCustomerRouter(customersRepo CustomerRepositoryI, countryFinderRepo countries.CountryFinder) CustomerRouter {
	return CustomerRouter{NewCustomerService(customersRepo, countryFinderRepo)}
}

func (r CustomerRouter) ListCustomers(c *gin.Context) {
	list, err := r.service.GetCategorizedCustomersList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal service error!"})
		return
	}

	c.JSON(http.StatusOK, list)
}
