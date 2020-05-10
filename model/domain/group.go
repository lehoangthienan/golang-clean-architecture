package domain

import (
	"github.com/jinzhu/gorm"
)

// Group model
type Group struct {
	Model
	Name      string `json:"name,omitempty"`
	CreatorID *UUID  `json:"creatorID,omitempty"`
}

// BeforeCreate prepare data before create data
func (o *Group) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", NewUUID())
	return nil
}
