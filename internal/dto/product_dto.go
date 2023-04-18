package dto

type ProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  string  `json:"category_id"`
}

type ProductResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryName string  `json:"category_name"`
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
