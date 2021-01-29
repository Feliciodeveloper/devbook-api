package models

import (
	"api/src/services"
	"gorm.io/gorm"
	"strings"
)

//entidade de usuarios para ser convertida no modelo relacional
type Posts struct {
	gorm.Model
	Title string `gorm:"size:100;not null;uniqueIndex:idx_post" json:"title,omitempty" validate:"required"`
	Content string `gorm:"not null" json:"content,omitempty" validate:"required"`
	Likes []Users `gorm:"many2many:posts_users" json:"likes,omitempty"`
	AuthorID uint `gorm:"uniqueIndex:idx_post" json:"-"`
	Author *Users `json:"author,omitempty"`
}
func(p *Posts)TableName()string{
	return "posts"
}

//metodos relacionados a validação da entidade
func(p *Posts)validate()error{
	if err := services.ValidateStruct(p);err != nil{
			return err
		}
	return nil
}
func(p *Posts)clear(){
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}
func(p *Posts)Treatment()error{
	if err := p.validate(); err != nil{
		return err
	}
	p.clear()
	return nil
}
