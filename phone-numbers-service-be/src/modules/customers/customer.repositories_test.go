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
		expectedCount int
	}

	tests := []test{
		{testName: "returns all customers", inFilters: []interface{}{""}, expectedCount: 41},
		{testName: "returns a customer by id", inFilters: []interface{}{10}, expectedCount: 1},
		{testName: "no customer with this phone", inFilters: []interface{}{"phone = 123"}, expectedCount: 0},
	}

	repo := NewCustomerRepository(db)

	for _, tc := range tests {
		customers, err := repo.GetCustomers(tc.inFilters...)
		if err != nil {
			t.Errorf("Test (%s) failed: unexpected error %+v", tc.testName, err)
		}

		if len(customers) != tc.expectedCount {
			t.Errorf("Test (%s) failed: expected %d customers, found %d customers", tc.testName, tc.expectedCount, len(customers))
		}
	}

}
