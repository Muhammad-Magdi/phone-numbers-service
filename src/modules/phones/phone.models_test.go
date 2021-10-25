package phones

import (
	"fmt"
	"testing"

	"phone-numbers-service.jpay.task/src/modules/countries"
)

func TestExtractCountryCode(t *testing.T) {
	type test struct {
		number   string
		code     string
		testName string
		errMsg   string
	}

	tests := []test{
		{number: "(201) 123456789", code: "201", testName: "correct number", errMsg: ""},
		{number: "(123)", code: "123", testName: "number consisted of only country code", errMsg: ""},
		{number: ")123(", code: "", testName: "reverted parantheses", errMsg: ERROR_INVALID_PHONE_NUMBER},
		{number: "(123", code: "", testName: "no closing parenthes", errMsg: ERROR_INVALID_PHONE_NUMBER},
		{number: "123)", code: "", testName: "no opening parenthes", errMsg: ERROR_INVALID_PHONE_NUMBER},
		{number: "", code: "", testName: "empty number", errMsg: ERROR_INVALID_PHONE_NUMBER},
	}

	for _, tc := range tests {
		returnedCode, err := extractCountryCode(tc.number)
		if (err == nil && tc.errMsg != "") || (err != nil && tc.errMsg == "") {
			t.Errorf("Test %s Failed: expected errMsg %s, thrown error = %s", tc.testName, tc.errMsg, err.Error())
		}

		if returnedCode != tc.code {
			t.Errorf("Test '%s' Failed: expected code = %s, returned code = %s", tc.testName, tc.code, returnedCode)
		}
	}
}

func TestIsValidPhoneNumber(t *testing.T) {
	type test struct {
		testName      string
		number        string
		countryRegExp string
		isValid       bool
	}

	tests := []test{
		{testName: "correct number", number: "(201) 123456789", countryRegExp: `\(201\)\ [0125]\d{8}$`, isValid: true},
		{testName: "valid number for another country", number: "(201) 123456789", countryRegExp: `\(212\)\ ?[5-9]\d{8}$`, isValid: false},
		{testName: "number consisted of only country code", number: "(201)", countryRegExp: `\(201\)\ [0125]\d{8}$`, isValid: false},
		{testName: "reverted parantheses", number: ")201(", countryRegExp: `\(201\)\ [0125]\d{8}$`, isValid: false},
		{testName: "no closing parenthes", number: "(201", countryRegExp: `\(201\)\ [0125]\d{8}$`, isValid: false},
		{testName: "no opening parenthes", number: "201)", countryRegExp: `\(201\)\ [0125]\d{8}$`, isValid: false},
		{testName: "empty number", number: "", countryRegExp: `\(201\)\ [0125]\d{8}$`, isValid: false},
	}

	for _, tc := range tests {
		isValid, err := isValidPhoneNumber(tc.number, tc.countryRegExp)
		if err != nil {
			t.Errorf("Test %s Failed: thrown error = %v", tc.testName, err)
		}

		if isValid != tc.isValid {
			t.Errorf("Test '%s' Failed: expected %t, found %t", tc.testName, tc.isValid, isValid)
		}
	}
}

func TestPhoneFactory(t *testing.T) {
	phoneFactory := PhoneFactory(mockedCountryRepository{})

	type test struct {
		testName string
		number   string
		country  *countries.Country
		isValid  bool
	}

	tests := []test{
		{testName: "valid egyptian number", number: "(201) 112346578", country: countries.NewCountry("Egypt", "201", `\(201\)\ [0125]\d{8}$`), isValid: true},
		{testName: "correct egyptian code, number missing one digit", number: "(201) 12346578", country: countries.NewCountry("Egypt", "201", `\(201\)\ [0125]\d{8}$`), isValid: false},
		{testName: "unknown country", number: "(205) 12346578", country: nil, isValid: false},
		{testName: "invalid code", number: "(205( 12346578", country: nil, isValid: false},
	}

	for _, tc := range tests {
		phone := phoneFactory(tc.number)

		if phone.IsValid() != tc.isValid {
			t.Errorf("Test '%s' failed: expected isValid = %t, found = %t", tc.testName, tc.isValid, phone.IsValid())
		}
		if phone.Number() != tc.number {
			t.Errorf("Test '%s' failed: expected isValid = %s, found = %s", tc.testName, tc.number, phone.Number())
		}

		if tc.country != nil {
			if phone.CountryName() != tc.country.Name() {
				t.Errorf("Test '%s' failed: expected countryName = %s, found = %s", tc.testName, tc.country.Name(), phone.CountryName())
			}
			if phone.CountryCode() != tc.country.Code() {
				t.Errorf("Test '%s' failed: expected countryCode = %s, found = %s", tc.testName, tc.country.Code(), phone.CountryCode())
			}
		} else {
			if phone.CountryName() != INVALID_COUNTRY {
				t.Errorf("Test '%s' failed: expected countryName = %s, found = %s", tc.testName, INVALID_COUNTRY, phone.CountryName())
			}
			if phone.CountryCode() != INVALID_COUNTRY {
				t.Errorf("Test '%s' failed: expected countryCode = %s, found = %s", tc.testName, INVALID_COUNTRY, phone.CountryCode())
			}
		}

	}
}

type mockedCountryRepository struct {
}

func (mockedCountryRepository) FindCountryByCode(code string) (*countries.Country, error) {
	if code == "201" {
		return countries.NewCountry("Egypt", "201", `\(201\)\ [0125]\d{8}$`), nil
	} else if code == "212" {
		return countries.NewCountry("Morocco", "212", `\(212\)\ ?[5-9]\d{8}$`), nil
	}
	return &countries.Country{}, fmt.Errorf("Invalid Country Code")
}

func assertPanic(t *testing.T, f func() interface{}) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
