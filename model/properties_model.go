package model

import "gorm.io/gorm"

type Properties struct {
	gorm.Model
	ID             uint          `gorm:"primaryKey"`
	Name           string        `json:"name"`
	Price          int           `json:"price"`
	Description    string        `json:"description"`
	
}

type CreatePropertiesRequest struct {

	
	Name           string        `json:"name"`
	Price          int           `json:"price"`
	Description    string        `json:"description"`
	
}
type CreatePropertiesResponse struct {
	
	ID             uint          `gorm:"primaryKey"`
	Name           string        `json:"name"`
	Price          int           `json:"price"`
	Description    string        `json:"description"`
	
}

type GetPropertiesResponse struct {
	
	ID             uint          `gorm:"primaryKey"`
	Name           string        `json:"name"`
	Price          int           `json:"price"`
	Description    string        `json:"description"`
	
}