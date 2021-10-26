package customers

type Customer struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func NewCustomer(id uint32, name string, phone string) Customer {
	c := Customer{}

	c.SetId(id)
	c.SetName(name)
	c.SetPhone(phone)

	return c
}

func (c *Customer) SetId(id uint32) {
	c.ID = id
}

func (c *Customer) SetName(name string) {
	c.Name = name
}

func (c *Customer) SetPhone(phone string) {
	c.Phone = phone
}
