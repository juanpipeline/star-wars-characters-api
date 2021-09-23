package models

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

type CharacterBase struct {
	Name      string  `json:"name"`
	Height    float64 `json:"height"`
	Mass      float64 `json:"mass"`
	HairColor string  `json:"hair_color"`
	SkinColor string  `json:"skin_color"`
	Masters   string  `json:"masters"`
	Species   string  `json:"species"`
}

type Character struct {
	CharacterBase
	gorm.Model
}

type CharacterCreate struct {
	CharacterBase
}

type CharacterResponse struct {
	CharacterBase
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}
