package countries

type CountryFinder interface {
	FindCountryByCode(code string) (*Country, error)
}

type CountryLister interface {
	GetCountryList() ([]*Country, error)
}
type CountryRepository struct {
}

func NewCountryRepository() CountryRepository {
	return CountryRepository{}
}

func (CountryRepository) GetCountryList() ([]*Country, error) {
	list := make([]*Country, 0, 5)

	list = append(list, NewCountry("Morocco", "212", `\(212\)\ ?[5-9]\d{8}$`))
	list = append(list, NewCountry("Cameron", "237", `\(237\)\ ?[2368]\d{7,8}$`))
	list = append(list, NewCountry("Ethiopia", "251", `\(251\)\ ?[1-59]\d{8}$`))
	list = append(list, NewCountry("Uganda", "256", `\(256\)\ ?\d{9}$`))
	list = append(list, NewCountry("Mozambique", "258", `\(258\)\ ?[28]\d{7,8}$`))

	return list, nil
}
