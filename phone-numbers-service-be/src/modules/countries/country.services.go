package countries

type CountryServiceI interface {
	ListCountries() []CountryDTO
}

type CountryService struct {
	repo CountryLister
}

func NewCountryService(repo CountryLister) CountryServiceI {
	return CountryService{repo}
}

type CountryDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// Retruns a list of all the countries, each with its Name and Code.
func (r CountryService) ListCountries() []CountryDTO {
	countryList, _ := r.repo.ListCountries()

	countryNames := make([]CountryDTO, 0, len(countryList))

	for _, c := range countryList {
		countryNames = append(countryNames, CountryDTO{Name: c.name, Code: c.code})
	}

	return countryNames
}
