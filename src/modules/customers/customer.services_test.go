package customers

import (
	"fmt"
	"testing"

	"phone-numbers-service.jpay.task/src/modules/countries"
	"phone-numbers-service.jpay.task/src/modules/phones"
)

func TestGetCategorizedCustomers(t *testing.T) {
	service := NewCustomerService(mockedCustomerRepository{}, mockedCountryRepository{})

	customers, err := service.GetCategorizedCustomersList()
	if err != nil {
		t.Errorf("Test failed: Unexpected error %+v", err)
	}

	if len(customers) != 5 {
		t.Errorf("Wrong customers count")
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
	if validCount != 2 {
		t.Errorf("Wrong valid customers count:\n expected: %d, found: %d", 2, validCount)
	}
	if validCountries != 3 {
		t.Errorf("Wrong valid customers count:\n expected: %d, found: %d", 2, validCount)
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
