package countries

import "testing"

func TestListCountries(t *testing.T) {
	repo := NewCountryRepository()

	countries, _ := repo.ListCountries()

	if len(countries) != 5 {
		t.Errorf("Testing ListCountries failed: expected = %d countries, found = %d", 5, len(countries))
	}
}

func TestFindCountryByCode(t *testing.T) {
	repo := NewCountryRepository()

	type test struct {
		testName            string
		inCode              string
		expectedCountryName string
		expectedErrMsg      string
	}

	tests := []test{
		{testName: "correct country code", inCode: "212", expectedCountryName: "Morocco", expectedErrMsg: ""},
		{testName: "wrong country code", inCode: "123", expectedCountryName: "", expectedErrMsg: ERROR_NO_SUCH_COUNTRY},
	}

	for _, tc := range tests {
		country, err := repo.FindCountryByCode(tc.inCode)

		if tc.expectedErrMsg == "" {
			if err != nil {
				t.Errorf("Test (%s) failed: Unexpected error %+v", tc.testName, err)
			}
			if country.name != tc.expectedCountryName {
				t.Errorf("Test (%s) failed: Expected countryName = %s, found = %s", tc.testName, tc.expectedCountryName, country.Name())
			}
		} else {
			if err == nil {
				t.Errorf("Test (%s) failed: Expected error %s", tc.testName, tc.expectedErrMsg)
			}
		}
	}
}
