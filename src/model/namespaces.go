package model

import "github.com/jinzhu/gorm"

type namespace struct {
	gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
}
