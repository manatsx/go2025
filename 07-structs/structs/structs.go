package structs

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
