package customers

import (
	"phone-numbers-service.jpay.task/src/modules/countries"
	"phone-numbers-service.jpay.task/src/modules/phones"
)

type CustomerServiceI interface {
	GetCategorizedCustomersList(filters CustomerFilters) ([]CustomerDTO, error)
}
type CustomerService struct {
	customerRepo CustomerGetterRepository
	countryRepo  countries.CountryFinder
}

func NewCustomerService(customerRepo CustomerGetterRepository, countryRepo countries.CountryFinder) CustomerService {
	service := CustomerService{customerRepo, countryRepo}
	return service
}

type CustomerDTO struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
	IsValid bool   `json:"is_valid"`
}

func NewCustomerDTO(repo countries.CountryFinder, customer Customer) CustomerDTO {
	phone := phones.PhoneFactory(repo)(customer.Phone)

	return CustomerDTO{
		ID:      customer.ID,
		Name:    customer.Name,
		Phone:   customer.Phone,
		Country: phone.CountryName(),
		IsValid: phone.IsValid(),
	}
}

// Used to filter the response of GetCategorizedCustomersList(filters)
//
// The response is filtered using both CountryName AND IsValid values, unless their values are nil.
type CustomerFilters struct {
	CountryName *string
	IsValid     *bool
}

// Returns a list of customerDTOs after applying the given filters.
func (s CustomerService) GetCategorizedCustomersList(filters CustomerFilters) ([]CustomerDTO, error) {
	customers, err := s.customerRepo.GetCustomers()
	// TODO: use app defined error instead of gorm errors
	if err != nil {
		return nil, err
	}

	customerDTOs := make([]CustomerDTO, 0, len(customers))
	for _, c := range customers {
		customer := NewCustomerDTO(s.countryRepo, c)

		if s.canAppendCustomer(customer, filters) {
			customerDTOs = append(customerDTOs, customer)
		}
	}

	return customerDTOs, nil
}

// Decides whether a given customer satisfies a given filter or not.
func (CustomerService) canAppendCustomer(customer CustomerDTO, filters CustomerFilters) bool {
	if (filters.CountryName != nil && customer.Country != *filters.CountryName) || (filters.IsValid != nil && customer.IsValid != *filters.IsValid) {
		return false
	}
	return true
}
