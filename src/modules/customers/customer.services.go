package customers

import (
	"phone-numbers-service.jpay.task/src/modules/countries"
	"phone-numbers-service.jpay.task/src/modules/phones"
)

type CustomerServiceI interface {
	GetCategorizedCustomersList() ([]CategorizedCustomerDTO, error)
}
type CustomerService struct {
	customerRepo CustomerRepositoryI
	countryRepo  countries.CountryRepositoryI
}

func NewCustomerService(customersRepo CustomerRepositoryI, countryRepo countries.CountryRepositoryI) CustomerService {
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

func NewCategorizedCustomer(repo countries.CountryRepositoryI, customer Customer) CategorizedCustomerDTO {
	phone := phones.PhoneFactory(repo)(customer.Phone)

	return CategorizedCustomerDTO{
		ID:      customer.ID,
		Name:    customer.Name,
		Phone:   customer.Phone,
		Country: phone.CountryName(),
		IsValid: phone.IsValid(),
	}
}
func (s CustomerService) GetCategorizedCustomersList() ([]CategorizedCustomerDTO, error) {
	customers, err := s.customerRepo.GetCustomers()
	// TODO: use app defined error instead of gorm errors
	if err != nil {
		return nil, err
	}

	categorizedCustomers := make([]CategorizedCustomerDTO, 0, len(customers))
	for _, c := range customers {
		categorizedCustomers = append(categorizedCustomers, NewCategorizedCustomer(s.countryRepo, c))
	}

	return categorizedCustomers, nil
}
