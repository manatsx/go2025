package structs

import "strconv"

type Product struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	TypeProduct Type     `json:"type"`
	Count       uint16   `json:"count"`
	Price       float64  `json:"price"`
	Tags        []string `json:"tags"`
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Car struct {
	ID       uint      `json:"id"`
	UserID   string    `json:"user_id"`
	Products []Product `json:"products"`
}

func (p Product) GetPrice() float64 {
	return float64(p.Count) * p.Price
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) AddTags(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

func (c *Car) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Car) Total() float64 {
	var total float64
	for _, c := range c.Products {
		total += c.GetPrice()
	}
	return total
}

func NewCar(userID uint) Car {
	return Car{UserID: strconv.Itoa(int(userID))}
}
