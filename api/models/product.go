package models

import (
	"fmt"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p Product) Display() string {
	return fmt.Sprintf("Product: %s - Price: %.2f", p.Name, p.Price)
}
