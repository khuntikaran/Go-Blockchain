package models

type Product struct {
	Productname string `json:"productname"`
	Price       string `json:"price"`
	Images      string `json:"images"`
	ProductType string `json:"producttype"`
	Sellername  string `json:"sellername"`
	Shopname    string `json:"shopname"`
	Description string `json:"description"`
}
