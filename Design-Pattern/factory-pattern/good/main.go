package main

type ProductFactory interface {
	Create() Product
}

type Product interface {
	GetName() string
	GetPrice() float64
}

type Shirt struct {
	name  string
	price float64
}

func (s *Shirt) GetName() string {
	return s.name
}

func (s *Shirt) GetPrice() float64 {
	return s.price
}

type ShirtFactory struct{}

func (sf *ShirtFactory) Create() Product {
	return &Shirt{name: "Basic shirt", price: 25.0}
}

type Pants struct {
	name  string
	price float64
}

func (p *Pants) GetName() string {
	return p.name
}

func (p *Pants) GetPrice() float64 {
	return p.price
}

type PantsFactory struct{}

func (pf *PantsFactory) Create() Product {
	return &Pants{name: "Basic pants", price: 35.0}
}
