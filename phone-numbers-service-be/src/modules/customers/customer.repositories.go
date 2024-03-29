package customers

import "gorm.io/gorm"

type CustomerGetterRepository interface {
	GetCustomers(filters ...interface{}) ([]Customer, error)
}

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return CustomerRepository{db}
}

// Returns a list of customers satisfying the given filters.
func (repo CustomerRepository) GetCustomers(filters ...interface{}) ([]Customer, error) {
	var customers []Customer
	err := repo.db.Find(&customers, filters...).Error
	// TODO: use app defined error instead of gorm errors
	if err != nil {
		return nil, err
	}

	return customers, nil
}
