package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	CategoryId  string
	ImageURL    string
}

func NewProduct(name, description, categoryId, imageURL string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CategoryId:  categoryId,
		ImageURL:    imageURL,
		Price:       price,
	}
}
