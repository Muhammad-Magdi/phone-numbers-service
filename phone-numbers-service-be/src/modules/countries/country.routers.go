package countries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CountryRouter struct {
	service CountryServiceI
}

func NewCountryRouter(repo CountryLister) CountryRouter {
	return CountryRouter{NewCountryService(repo)}
}

func (r CountryRouter) ListCountries(c *gin.Context) {
	list := r.service.ListCountries()

	c.JSON(http.StatusOK, list)
}
