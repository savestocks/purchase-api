package domain

import (
    "time"
)

//Purchase represents a model
type Purchase struct {
    ID string `json:"id"`
    CreatedAt time.Time `json:"dateCreated,omitempty"`
    UpdatedAt time.Time `json:"dateUpdated,omitempty"`
    ItemID string `json:"itemId"`
	Qtd    float32     `json:"qtd"`
	Price  int32     `json:"price"`
    MarketID string    `json:"marketId"`    
}