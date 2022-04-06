package book

import (
	"encoding/json"
)

// Struct data Post Books
type BookInput struct {
	Title string      `json:"title" binding:"required"`        //Validation not null / empty
	Price json.Number `json:"price" binding:"required,number"` //Validation not null / empty & json.number convert string to int
	// SubTitle string      `json:"sub_title"`                       //if data was not same
}
