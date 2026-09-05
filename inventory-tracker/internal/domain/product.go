package domain

import "time"

type Product struct {
	ID          string
	SKU         string
	Name        string
	Description *string
	Unit        string
	MinStock    int32
	OnHand      int32
	IsActive    bool
	StockState  string
	Version     int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}