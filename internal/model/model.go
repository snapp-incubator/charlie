package model

import (
	"gorm.io/gorm"
)

// Base is the first model of our models.
type Base struct {
	gorm.Model
}
