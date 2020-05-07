package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// TokenUser model
type TokenUser struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

// User model
type User struct {
	Model
	Name     string `json:"name,omitempty"`
	UserName string `json:"userName,omitempty"`
	PassWord string `json:"-"`
	Role     string `json:"role,omitempty"`
}

// BeforeCreate prepare data before create data
func (o *User) BeforeCreate(scope *gorm.Scope) error {
	var (
		hash []byte
		err  error
	)

	if o.PassWord != "" {
		hash, err = bcrypt.GenerateFromPassword([]byte(o.PassWord), bcrypt.DefaultCost)
	} else {
		hash, err = bcrypt.GenerateFromPassword([]byte((uuid.NewV4()).String()), bcrypt.DefaultCost)
	}

	if err != nil {
		return err
	}

	scope.SetColumn("PassWord", string(hash))
	scope.SetColumn("ID", NewUUID())
	return nil
}

// BeforeUpdate prepare data before update date
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	if u.PassWord != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.PassWord), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.PassWord = string(hash)
	}

	return nil
}

// ComparePassword func
func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(password))
	if err != nil {
		return false
	}
	return true
}
