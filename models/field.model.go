package models

import "gorm.io/gorm"

type NegiField struct {
	gorm.Model
	FieldName    string `json:"field_name,omitempty"`
	GroupName    string `json:"group_name,omitempty"`
	Active       bool   `json:"active,omitempty"`
	Address      string `json:"address,omitempty"`
	TaskCalEvent []TaskCalEvent
}

func init() {
	db.AutoMigrate(NegiField{})
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

	if err = db.Model(&F).First(F.TaskCalEvent).Error; err != nil {
		return
	}
	return
}

func (F *NegiField) GetAllNegiField() (fs []NegiField, err error) {
	if err = db.Find(&fs).Error; err != nil {
		return
	}
	return
}
