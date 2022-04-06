package book

import (
	"encoding/json"
)

// Struct data Post Books
type BookRequest struct {
	Title       string      `json:"title" binding:"required"`           //Validation not null / empty
	Price       json.Number `json:"price" binding:"required,number"`    //Validation not null / empty & json.number convert string to int
	Description string      `json:"description" binding:"required"`     //Validation not null / empty
	Rating      json.Number `json:"rating" binding:"required,number"`   //Validation not null / empty
	Discount    json.Number `json:"discount" binding:"required,number"` //Validation not null / empty

}
