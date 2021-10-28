package countries

import "errors"

type CountryFinder interface {
	FindCountryByCode(code string) (*Country, error)
}

type CountryLister interface {
	ListCountries() ([]*Country, error)
}
type CountryRepository struct {
}

func NewCountryRepository() CountryRepository {
	return CountryRepository{}
}

// Retruns a list of all the supported countries.
//
// Currently, there're only 5 fixed countries that are supported.
func (CountryRepository) ListCountries() ([]*Country, error) {
	list := make([]*Country, 0, 5)

	list = append(list, NewCountry("Morocco", "212", `\(212\)\ ?[5-9]\d{8}$`))
	list = append(list, NewCountry("Cameron", "237", `\(237\)\ ?[2368]\d{7,8}$`))
	list = append(list, NewCountry("Ethiopia", "251", `\(251\)\ ?[1-59]\d{8}$`))
	list = append(list, NewCountry("Uganda", "256", `\(256\)\ ?\d{9}$`))
	list = append(list, NewCountry("Mozambique", "258", `\(258\)\ ?[28]\d{7,8}$`))

	return list, nil
}

// Takes a country phone code and returns a pointer to that country.
//
// Returns an `ERROR_NO_SUCH_COUNTRY`, if there isn't a country with the given code.
func (r CountryRepository) FindCountryByCode(code string) (*Country, error) {
	list, _ := r.ListCountries()

	for _, c := range list {
		if c.Code() == code {
			return c, nil
		}
	}

	return nil, errors.New(ERROR_NO_SUCH_COUNTRY)
}
