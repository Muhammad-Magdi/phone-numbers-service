package customers

import (
	"phone-numbers-service.jpay.task/src/modules/countries"
	"phone-numbers-service.jpay.task/src/modules/phones"
)

type CustomerServiceI interface {
	GetCategorizedCustomersList(filters CustomerFilters) ([]CategorizedCustomerDTO, error)
}
type CustomerService struct {
	customerRepo CustomerRepositoryI
	countryRepo  countries.CountryFinder
}

func NewCustomerService(customersRepo CustomerRepositoryI, countryRepo countries.CountryFinder) CustomerService {
	service := CustomerService{customersRepo, countryRepo}
	return service
}

type CategorizedCustomerDTO struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
	IsValid bool   `json:"is_valid"`
}
type CustomerFilters struct {
	CountryName *string
	IsValid     *bool
}

func (CustomerService) canAppendCustomer(customer CategorizedCustomerDTO, filters CustomerFilters) bool {
	if (filters.CountryName != nil && customer.Country != *filters.CountryName) || (filters.IsValid != nil && customer.IsValid != *filters.IsValid) {
		return false
	}
	return true
}

func NewCategorizedCustomer(repo countries.CountryFinder, customer Customer) CategorizedCustomerDTO {
	phone := phones.PhoneFactory(repo)(customer.Phone)

	return CategorizedCustomerDTO{
		ID:      customer.ID,
		Name:    customer.Name,
		Phone:   customer.Phone,
		Country: phone.CountryName(),
		IsValid: phone.IsValid(),
	}
}
func (s CustomerService) GetCategorizedCustomersList(filters CustomerFilters) ([]CategorizedCustomerDTO, error) {
	customers, err := s.customerRepo.GetCustomers()
	// TODO: use app defined error instead of gorm errors
	if err != nil {
		return nil, err
	}

	categorizedCustomers := make([]CategorizedCustomerDTO, 0, len(customers))
	for _, c := range customers {
		customer := NewCategorizedCustomer(s.countryRepo, c)

		if s.canAppendCustomer(customer, filters) {
			categorizedCustomers = append(categorizedCustomers, customer)
		}
	}

	return categorizedCustomers, nil
}
