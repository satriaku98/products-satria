package model

import "time"

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique"`
	Price       uint      `json:"price"`
	Description string    `json:"description"`
	Quantity    uint      `json:"quantity"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
