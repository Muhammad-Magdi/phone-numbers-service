package countries

import "testing"

func TestGetCountryList(t *testing.T) {
	repo := NewCountryRepository()

	countries, _ := repo.GetCountryList()

	if len(countries) != 5 {
		t.Errorf("Testing GetCountryList failed: expected = %d countries, found = %d", 5, len(countries))
	}
}

func TestFindCountryByCode(t *testing.T) {
	repo := NewCountryRepository()

	type test struct {
		testName    string
		inCode      string
		countryName string
		err         string
	}

	tests := []test{
		{testName: "correct country code", inCode: "212", countryName: "Morocco", err: ""},
		{testName: "wrong country code", inCode: "123", countryName: "", err: ERROR_NO_SUCH_COUNTRY},
	}

	for _, tc := range tests {
		country, err := repo.FindCountryByCode(tc.inCode)

		if tc.err == "" {
			if err != nil {
				t.Errorf("Test (%s) failed: Unexpected error %+v", tc.testName, err)
			}
			if country.name != tc.countryName {
				t.Errorf("Test (%s) failed: Expected countryName = %s, found = %s", tc.testName, tc.countryName, country.Name())
			}
		} else {
			if err == nil {
				t.Errorf("Test (%s) failed: Expected error %s", tc.testName, tc.err)
			}
		}
	}
}
