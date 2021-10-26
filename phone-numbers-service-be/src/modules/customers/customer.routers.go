package customers

import (
	"net/http"
	"strconv"

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
	countryName := c.Query("country")
	isValid := c.Query("is_valid")

	filters := CustomerFilters{}

	if countryName != "" {
		filters.CountryName = &countryName
	}

	boolIsValid, err := strconv.ParseBool(isValid)
	if err == nil {
		filters.IsValid = &boolIsValid
	}

	list, err := r.service.GetCategorizedCustomersList(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal service error!"})
		return
	}

	c.JSON(http.StatusOK, list)
}
