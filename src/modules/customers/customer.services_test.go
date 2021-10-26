package customers

import (
	"fmt"
	"testing"

	"phone-numbers-service.jpay.task/src/modules/countries"
	"phone-numbers-service.jpay.task/src/modules/phones"
)

func TestGetCategorizedCustomers(t *testing.T) {
	service := NewCustomerService(mockedCustomerRepository{}, mockedCountryFinder{})

	type test struct {
		testName               string
		inFilters              CustomerFilters
		expectedCount          int
		expectedValidCount     int
		expectedValidCountries int
	}

	trueVar := true
	falseVar := false

	morocco := "Morocco"
	tunis := "Tunis"

	tests := []test{
		{testName: "No Filters", inFilters: CustomerFilters{}, expectedCount: 5, expectedValidCount: 2, expectedValidCountries: 3},
		{testName: "IsValid = true", inFilters: CustomerFilters{IsValid: &trueVar}, expectedCount: 2, expectedValidCount: 2, expectedValidCountries: 2},
		{testName: "IsValid = false", inFilters: CustomerFilters{IsValid: &falseVar}, expectedCount: 3, expectedValidCount: 0, expectedValidCountries: 1},
		{testName: "Country = Morocco", inFilters: CustomerFilters{CountryName: &morocco}, expectedCount: 2, expectedValidCount: 1, expectedValidCountries: 2},
		{testName: "Country = Tunis", inFilters: CustomerFilters{CountryName: &tunis}, expectedCount: 0, expectedValidCount: 0, expectedValidCountries: 0},
		{testName: "Country = Morocco & IsValid = true", inFilters: CustomerFilters{CountryName: &morocco, IsValid: &trueVar}, expectedCount: 1, expectedValidCount: 1, expectedValidCountries: 1},
		{testName: "Country = Morocco & IsValid = false", inFilters: CustomerFilters{CountryName: &morocco, IsValid: &falseVar}, expectedCount: 1, expectedValidCount: 0, expectedValidCountries: 1},
	}

	for _, tc := range tests {
		customers, err := service.GetCategorizedCustomersList(tc.inFilters)
		if err != nil {
			t.Errorf("Test (%s) failed: Unexpected error %+v", tc.testName, err)
		}

		fmt.Printf("%+v\n", customers)
		if len(customers) != tc.expectedCount {
			t.Errorf("Test (%s) failed: Expected totalCount = %d, found = %d", tc.testName, tc.expectedCount, len(customers))
		}

		validCount, validCountries := 0, 0
		for _, customer := range customers {
			if customer.IsValid {
				validCount++
			}
			if customer.Country != phones.INVALID_COUNTRY {
				validCountries++
			}
		}
		if validCount != tc.expectedValidCount {
			t.Errorf("Test (%s) failed: Expected validCount = %d, found = %d", tc.testName, tc.expectedValidCount, validCount)
		}
		if validCountries != tc.expectedValidCountries {
			t.Errorf("Test (%s) failed: Expected validCountriesCount = %d, found = %d", tc.testName, tc.expectedValidCountries, validCountries)
		}
	}

}

type mockedCustomerRepository struct {
}

func (mockedCustomerRepository) GetCustomers(filters ...interface{}) ([]Customer, error) {
	customers := make([]Customer, 5)

	customers[0] = NewCustomer(1, "Valid Egyptian number", "(201) 123456789")
	customers[1] = NewCustomer(2, "Valid Moroccan number", "(212) 512345678")
	customers[2] = NewCustomer(3, "Valid Moroccan code", "(212) xxo")
	customers[3] = NewCustomer(4, "Customer 4", "(2124679")
	customers[4] = NewCustomer(5, "Customer 5", ")212( 1234154")

	return customers, nil
}

type mockedCountryFinder struct {
}

func (mockedCountryFinder) FindCountryByCode(code string) (*countries.Country, error) {
	if code == "201" {
		return countries.NewCountry("Egypt", "201", `\(201\)\ [0125]\d{8}$`), nil
	} else if code == "212" {
		return countries.NewCountry("Morocco", "212", `\(212\)\ ?[5-9]\d{8}$`), nil
	}
	return &countries.Country{}, fmt.Errorf("Invalid Country Code")
}
