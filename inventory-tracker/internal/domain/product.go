package domain

import (
	"time"

	"github.com/google/uuid"
)

type StockState string

const (
	StockStateOutOfStock StockState = "OUT_OF_STOCK"
	StockStateLow        StockState = "LOW"
	StockStateHealthy    StockState = "HEALTHY"
)

type Product struct {
	ID          uuid.UUID  `json:"id"`
	SKU         string     `json:"sku"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Unit        string     `json:"unit"`
	MinStock    int32      `json:"min_stock"`
	OnHand      int32      `json:"on_hand"`
	IsActive    bool       `json:"is_active"`
	StockState  StockState `json:"stock_state"`
	Version     int32      `json:"version"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}