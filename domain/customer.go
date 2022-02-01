package domain

import "github.com/joaopcm/go-banking-api/errs"

type Customer struct {
	Id string `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
	DateOfBirth string `json:"date_of_birth" xml:"dateofbirth"`
	Status string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
