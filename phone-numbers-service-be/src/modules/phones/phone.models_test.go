package phones

import (
	"fmt"
	"testing"

	"phone-numbers-service.jpay.task/src/modules/countries"
)

func TestExtractCountryCode(t *testing.T) {
	type test struct {
		testName       string
		inNumber       string
		inCode         string
		expectedErrMsg string
	}

	tests := []test{
		{inNumber: "(201) 123456789", inCode: "201", testName: "correct number", expectedErrMsg: ""},
		{inNumber: "(123)", inCode: "123", testName: "number consisted of only country code", expectedErrMsg: ""},
		{inNumber: ")123(", inCode: "", testName: "reverted parantheses", expectedErrMsg: ERROR_INVALID_PHONE_NUMBER},
		{inNumber: "(123", inCode: "", testName: "no closing parenthes", expectedErrMsg: ERROR_INVALID_PHONE_NUMBER},
		{inNumber: "123)", inCode: "", testName: "no opening parenthes", expectedErrMsg: ERROR_INVALID_PHONE_NUMBER},
		{inNumber: "", inCode: "", testName: "empty number", expectedErrMsg: ERROR_INVALID_PHONE_NUMBER},
	}

	for _, tc := range tests {
		returnedCode, err := extractCountryCode(tc.inNumber)
		if (err == nil && tc.expectedErrMsg != "") || (err != nil && tc.expectedErrMsg == "") {
			t.Errorf("Test (%s) failed: expected errMsg %s, thrown = %s", tc.testName, tc.expectedErrMsg, err.Error())
		}

		if returnedCode != tc.inCode {
			t.Errorf("Test (%s) failed: expected code = %s, returned = %s", tc.testName, tc.inCode, returnedCode)
		}
	}
}

func TestIsValidPhoneNumber(t *testing.T) {
	type test struct {
		testName        string
		inNumber        string
		inCountryRegExp string
		expectedIsValid bool
	}

	tests := []test{
		{testName: "correct number", inNumber: "(201) 123456789", inCountryRegExp: `\(201\)\ [0125]\d{8}$`, expectedIsValid: true},
		{testName: "valid number for another country", inNumber: "(201) 123456789", inCountryRegExp: `\(212\)\ ?[5-9]\d{8}$`, expectedIsValid: false},
		{testName: "number consisted of only country code", inNumber: "(201)", inCountryRegExp: `\(201\)\ [0125]\d{8}$`, expectedIsValid: false},
		{testName: "reverted parantheses", inNumber: ")201(", inCountryRegExp: `\(201\)\ [0125]\d{8}$`, expectedIsValid: false},
		{testName: "no closing parenthes", inNumber: "(201", inCountryRegExp: `\(201\)\ [0125]\d{8}$`, expectedIsValid: false},
		{testName: "no opening parenthes", inNumber: "201)", inCountryRegExp: `\(201\)\ [0125]\d{8}$`, expectedIsValid: false},
		{testName: "empty number", inNumber: "", inCountryRegExp: `\(201\)\ [0125]\d{8}$`, expectedIsValid: false},
	}

	for _, tc := range tests {
		isValid, err := isValidPhoneNumber(tc.inNumber, tc.inCountryRegExp)
		if err != nil {
			t.Errorf("Test (%s) failed: unexpected error = %v", tc.testName, err)
		}

		if isValid != tc.expectedIsValid {
			t.Errorf("Test (%s) failed: expected isValid = %t, found = %t", tc.testName, tc.expectedIsValid, isValid)
		}
	}
}

func TestPhoneFactory(t *testing.T) {
	phoneFactory := PhoneFactory(mockedCountryRepository{})

	type test struct {
		testName        string
		inNumber        string
		inCountry       *countries.Country
		expectedIsValid bool
	}

	tests := []test{
		{testName: "valid egyptian number", inNumber: "(201) 112346578", inCountry: countries.NewCountry("Egypt", "201", `\(201\)\ [0125]\d{8}$`), expectedIsValid: true},
		{testName: "correct egyptian code, number missing one digit", inNumber: "(201) 12346578", inCountry: countries.NewCountry("Egypt", "201", `\(201\)\ [0125]\d{8}$`), expectedIsValid: false},
		{testName: "unknown country", inNumber: "(205) 12346578", inCountry: nil, expectedIsValid: false},
		{testName: "invalid code", inNumber: "(205( 12346578", inCountry: nil, expectedIsValid: false},
	}

	for _, tc := range tests {
		phone := phoneFactory(tc.inNumber)

		if phone.IsValid() != tc.expectedIsValid {
			t.Errorf("Test (%s) failed: expected isValid = %t, found = %t", tc.testName, tc.expectedIsValid, phone.IsValid())
		}
		if phone.Number() != tc.inNumber {
			t.Errorf("Test (%s) failed: expected number = %s, found = %s", tc.testName, tc.inNumber, phone.Number())
		}

		if tc.inCountry != nil {
			if phone.CountryName() != tc.inCountry.Name() {
				t.Errorf("Test (%s) failed: expected countryName = %s, found = %s", tc.testName, tc.inCountry.Name(), phone.CountryName())
			}
			if phone.CountryCode() != tc.inCountry.Code() {
				t.Errorf("Test (%s) failed: expected countryCode = %s, found = %s", tc.testName, tc.inCountry.Code(), phone.CountryCode())
			}
		} else {
			if phone.CountryName() != INVALID_COUNTRY {
				t.Errorf("Test (%s) failed: expected countryName = %s, found = %s", tc.testName, INVALID_COUNTRY, phone.CountryName())
			}
			if phone.CountryCode() != INVALID_COUNTRY {
				t.Errorf("Test (%s) failed: expected countryCode = %s, found = %s", tc.testName, INVALID_COUNTRY, phone.CountryCode())
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
