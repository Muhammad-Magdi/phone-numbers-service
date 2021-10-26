package countries

type CountryServiceI interface {
	ListCountries() []CountryNameDTO
}

type CountryService struct {
	repo CountryLister
}

func NewCountryService(repo CountryLister) CountryServiceI {
	return CountryService{repo}
}

type CountryNameDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (r CountryService) ListCountries() []CountryNameDTO {
	countryList, _ := r.repo.GetCountryList()

	countryNames := make([]CountryNameDTO, 0, len(countryList))

	for _, c := range countryList {
		countryNames = append(countryNames, CountryNameDTO{Name: c.name, Code: c.code})
	}

	return countryNames
}
