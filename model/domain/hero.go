package domain

import (
	"github.com/jinzhu/gorm"
)

// Hero model
type Hero struct {
	Model
	Name      string `json:"name,omitempty"`
	Power     string `json:"power,omitempty"`
	CreatorID *UUID  `json:"creatorID,omitempty"`
}

// BeforeCreate prepare data before create data
func (o *Hero) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", NewUUID())
	return nil
}
