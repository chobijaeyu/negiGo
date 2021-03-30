package models

import "gorm.io/gorm"

type TaskTitleOption struct {
	gorm.Model
	Title string `json:"title,omitempty"`
}

func init() {
	db.AutoMigrate(TaskTitleOption{})
}

func (T *TaskTitleOption) CreateTaskTitleOption() (err error) {
	if err = db.Create(&T).Error; err != nil {
		return
	}
	return
}

func (T *TaskTitleOption) UpdatesTaskTitleOption() (err error) {

	if err = db.Updates(&T).Error; err != nil {
		return
	}
	return
}

func (T *TaskTitleOption) DeleteTaskTitleOption() (err error) {
	if err = db.Delete(&T).Error; err != nil {
		return
	}
	return
}

func (T *TaskTitleOption) GetAllTaskTitleOption() (ts []TaskTitleOption, err error) {
	if err = db.Find(&ts).Error; err != nil {
		return
	}
	return
}
