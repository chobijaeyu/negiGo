package models

import (
	"github.com/jinzhu/gorm"
)

type NegiField struct {
	gorm.Model
	FieldName string `json:"field_name,omitempty"`
	GroupName string `json:"group_name,omitempty"`
}

func init() {

}

func (F *NegiField) CreateNegiField() (err error) {

	if err = db.Create(&F).Error; err != nil {
		return
	}
	return
}

func (F *NegiField) UpdatesNegiField() (err error) {

	// structs.Map(&F)
	if err = db.Updates(&F).Error; err != nil {
		return
	}
	return
}

func (F *NegiField) DeleteNegiField() (err error) {

	if err = db.Delete(&F).Error; err != nil {
		return
	}
	return
}

func (F *NegiField) GetOneNegiField() (err error) {

	if err = db.First(F.ID).Error; err != nil {
		return
	}
	return
}
