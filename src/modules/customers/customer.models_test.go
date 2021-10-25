package customers

import "testing"

func TestNewCustomer(t *testing.T) {
	type test struct {
		testName string
		inId     uint32
		inName   string
		inPhone  string
	}

	tests := []test{
		{testName: "should create customer with values the same as input", inId: 1, inName: "new customer", inPhone: "(201) 123456789"},
	}

	for _, tc := range tests {
		customer := NewCustomer(tc.inId, tc.inName, tc.inPhone)

		if customer.ID != tc.inId || customer.Name != tc.inName || customer.Phone != tc.inPhone {
			t.Errorf("Test '%s' failed: expected id = %d, name = %s, phone = %s, found customer = %+v", tc.testName, tc.inId, tc.inName, tc.inPhone, customer)
		}
	}
}
