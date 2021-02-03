package models

import (
	"api/src/safety"
	"api/src/services"
	"gorm.io/gorm"
	"strings"
)

//entidade de usuarios para ser convertida no modelo relacional
type Users struct {
	gorm.Model
	Name     string  `gorm:"size:50;not null" json:"name,omitempty" validate:"required"`
	Nick     string  `gorm:"size:50;unique;not null" json:"nick,omitempty" validate:"required"`
	Email    string  `gorm:"size:50;unique;not null" json:"email,omitempty" validate:"required,email"`
	Password string  `gorm:"not null" validate:"required"`
	Friends  []Users `gorm:"many2many:users_friends;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"friends,omitempty"`
	Posts    []Posts `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"posts,omitempty"`
}

func (u *Users) TableName() string {
	return "users"
}

//estrutura auxiliar a entidade
type Password struct {
	Old string
	New string
}

//metodos relacionados a validação da entidade
func (u *Users) validate(param string) error {
	if param == "create" {
		if err := services.ValidateStruct(u); err != nil {
			return err
		}
	}
	if u.Email != "" {
		if err := services.ValidateEmail(u.Email); err != nil {
			return err
		}
	}
	return nil
}
func (u *Users) clear() error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	PasswordHAsh, err := safety.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(PasswordHAsh)
	return nil
}
func (u *Users) Treatment(param string) error {
	if err := u.validate(param); err != nil {
		return err
	}
	if err := u.clear(); err != nil {
		return err
	}
	return nil
}
