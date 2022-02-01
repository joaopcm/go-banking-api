package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{ Id: "1", Name: "João", City: "São Paulo", Zipcode: "00000000", DateOfBirth: "2002-09-06", Status: "1" },
		{ Id: "2", Name: "John Doe", City: "New York", Zipcode: "00000000", DateOfBirth: "1995-05-23", Status: "1" },
	}

	return CustomerRepositoryStub{customers}
}