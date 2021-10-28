package phones

import (
	"errors"
	"regexp"
	"strings"

	"phone-numbers-service.jpay.task/src/modules/countries"
)

type Phone struct {
	number  string
	country *countries.Country
	isValid bool
}

func PhoneFactory(countryRepo countries.CountryFinder) func(string) Phone {
	return func(number string) Phone {
		phone := Phone{}

		phone.setNumber(number)
		phone.setCountry(countryRepo)
		phone.validate()

		return phone
	}
}

func (p *Phone) setNumber(number string) {
	p.number = number
}

// Extracts the countryCode from the phone number, uses the countryFinder to find the corresponding country then sets it.
func (p *Phone) setCountry(countryRepo countries.CountryFinder) {
	countryCode, err := extractCountryCode(p.Number())
	if err != nil {
		p.country = nil
		return
	}

	country, err := countryRepo.FindCountryByCode(countryCode)
	if err != nil {
		p.country = nil
		return
	}

	p.country = country
}

// Sets the isValid field of the phone depending on its number and country regexp.
func (p *Phone) validate() {
	if p.country == nil {
		p.isValid = false
		return
	}

	isValid, err := regexp.MatchString(p.country.PhoneRegExp(), p.Number())
	if err != nil {
		isValid = false
		return
	}

	p.isValid = isValid
}

func (p Phone) Number() string {
	return p.number
}

// Returns the country code that this phone (number) exists in.
//
// If the phone.Number() is in an unknown or invalid country, it returns `INVALID_COUNTRY`.
func (p Phone) CountryCode() string {
	if p.country == nil {
		return INVALID_COUNTRY
	}

	return p.country.Code()
}

// Returns the country name that this phone (number) exists in.
//
// If the phone.Number() is in an unknown or invalid country, it returns `INVALID_COUNTRY`.
func (p Phone) CountryName() string {
	if p.country == nil {
		return INVALID_COUNTRY
	}

	return p.country.Name()
}

func (p Phone) IsValid() bool {
	return p.isValid
}

// Extracts the country code from a phone number in the form: `(COUNTRYCODE) REST_OF_THE_NUMBER`
func extractCountryCode(number string) (string, error) {
	number = strings.TrimSpace(number)
	if len(number) < 3 || number[0] != '(' {
		return "", errors.New(ERROR_INVALID_PHONE_NUMBER)
	}

	var codeBuilder strings.Builder
	i := 1
	for ; i < len(number) && number[i] != ')'; i++ {
		codeBuilder.WriteByte(number[i])
	}

	if i == len(number) {
		return "", errors.New(ERROR_INVALID_PHONE_NUMBER)
	}

	return codeBuilder.String(), nil
}

// Decides whether a given phone number matches the given country regular expression.
func isValidPhoneNumber(number string, countryRegExp string) (bool, error) {
	return regexp.MatchString(countryRegExp, number)
}
