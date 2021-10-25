package countries

import "testing"

func TestGetCountryList(t *testing.T) {
	repo := NewCountryRepository()

	countries := repo.GetCountryList()

	if len(countries) != 5 {
		t.Errorf("Testing GetCountryList failed: expected = %d countries, found = %d", 5, len(countries))
	}
}
