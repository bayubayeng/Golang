package book

import (
	"encoding/json"
	"time"
)

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      int         `json :"rating" binding:"required"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdateAt    time.Time   `json:"updateAt"`
}
