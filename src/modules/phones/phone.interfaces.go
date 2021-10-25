package phones

import "phone-numbers-service.jpay.task/src/modules/countries"

type CountryRepository interface {
	FindCountryByCode(code string) (*countries.Country, error)
}
