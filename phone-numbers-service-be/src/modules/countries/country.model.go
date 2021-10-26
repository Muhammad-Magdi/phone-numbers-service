package countries

type Country struct {
	name        string
	code        string
	phoneRegExp string
}

func NewCountry(name string, code string, phoneRegExp string) *Country {
	return &Country{name: name, code: code, phoneRegExp: phoneRegExp}
}

// ====== Setters ======

func (c *Country) SetName(name string) {
	c.name = name
}

func (c *Country) SetCode(code string) {
	c.code = code
}

func (c *Country) SetPhoneRegExp(phoneRegExp string) {
	c.phoneRegExp = phoneRegExp
}

// ====== Getters ======

func (c Country) Name() string {
	return c.name
}

func (c Country) Code() string {
	return c.code
}

func (c Country) PhoneRegExp() string {
	return c.phoneRegExp
}
