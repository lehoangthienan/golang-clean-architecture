package domain

import (
	"github.com/jinzhu/gorm"
)

// GroupHero struct contain info of a GroupHero
type GroupHero struct {
	Model
	GroupID *UUID `json:"groupID,omitempty"`
	HeroID  *UUID `json:"heroID,omitempty"`

	// Preload
	Group *Group `gorm:"PRELOAD:false" json:"group,omitempty"`
	Hero  *Hero  `gorm:"PRELOAD:false" json:"hero,omitempty"`
}

// BeforeCreate prepare data before create data
func (o *GroupHero) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", NewUUID())
	return nil
}
