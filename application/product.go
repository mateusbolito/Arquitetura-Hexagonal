package application

import "errors"

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	disable() error
	GetId() string
	GetName() string
	getStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	Enabled  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

// func (p *Product) IsValid() (bool, error) {

// }
func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = "enabled"
		return nil
	}
	return errors.New("the price must be greater than zero enable the product")
}

// func (p *Product) Disable() error {

// }

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
