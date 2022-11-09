package domain

type Product struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Count int `json:"count"`
	Price float64 `json:"price"`
}