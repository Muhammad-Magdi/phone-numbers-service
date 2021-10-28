package countries

type Country struct {
	name        string
	code        string
	phoneRegExp string
}

// Returns a pointer to a new country with the specified parameters
func NewCountry(name string, code string, phoneRegExp string) *Country {
	country := Country{}
	country.SetCode(code)
	country.SetName(name)
	country.SetPhoneRegExp(phoneRegExp)
	return &country
}

func (c *Country) SetName(name string) {
	c.name = name
}

func (c *Country) SetCode(code string) {
	c.code = code
}

func (c *Country) SetPhoneRegExp(phoneRegExp string) {
	c.phoneRegExp = phoneRegExp
}

func (c Country) Name() string {
	return c.name
}

func (c Country) Code() string {
	return c.code
}

func (c Country) PhoneRegExp() string {
	return c.phoneRegExp
}
