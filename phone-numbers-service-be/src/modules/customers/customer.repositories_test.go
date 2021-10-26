package customers

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetCustomers(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../../../db/sample.db"), &gorm.Config{})
	if err != nil {
		t.Errorf("Test GetCustomers failed: couldn't open DB %+v", err)
	}

	type test struct {
		testName      string
		inFilters     []interface{}
		returnedCount int
	}

	tests := []test{
		{testName: "returns all customers", inFilters: []interface{}{""}, returnedCount: 41},
		{testName: "returns a customer by id", inFilters: []interface{}{10}, returnedCount: 1},
		{testName: "no customer with this phone", inFilters: []interface{}{"phone = 123"}, returnedCount: 0},
	}

	repo := NewCustomerRepository(db)

	for _, tc := range tests {
		customers, err := repo.GetCustomers(tc.inFilters...)
		if err != nil {
			t.Errorf("Test GetCustomers failed: unexpected error %+v", err)
		}

		if len(customers) != tc.returnedCount {
			t.Errorf("Test GetCustomers failed: expected %d customers, found %d customers", tc.returnedCount, len(customers))
		}
	}

}
