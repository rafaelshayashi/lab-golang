package model

type Customer struct {
	Name        string
	Address     string
	DateOfBirth string
}

func NewCustomer(name string, address string, dataOfBirth string) *Customer {
	return &Customer{
		Name:        name,
		Address:     address,
		DateOfBirth: dataOfBirth,
	}
}

func (c *Customer) ToSlice() []string {
	return []string{c.Name, c.Address, c.DateOfBirth}
}
