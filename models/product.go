package models

type Product struct {
	ID          int    `bun:"id,pk,autoincrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
}
