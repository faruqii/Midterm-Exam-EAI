package dto

import "time"

type ProductRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  string    `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductUpdateRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  string    `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ProductResponseList struct {
	Products []ProductResponse `json:"products"`
}

type ProductResponseByID struct {
	Product ProductResponse `json:"product"`
}

type ProductResponseByCategoryName struct {
	Products []ProductResponse `json:"products"`
}

type ProductResponseByName struct {
	Product ProductResponse `json:"product"`
}

type ProductResponseByCategoryNameList struct {
	Products []ProductResponseByCategoryName `json:"products"`
}

type CategoriesResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
