package models

type Response struct {
    TotalItems		int	`json:"total_items"`
    TotalCategories int	`json:"total_categories"`
    TotalPrice		int `json:"total_price"`
}
