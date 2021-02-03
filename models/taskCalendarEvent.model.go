package models

import (
	"gorm.io/gorm"
)

type TaskCalEvent struct {
	gorm.Model
	Start     string `json:"start" firestore:"start"`
	End       string `json:"end" firestore:"end"`
	Title     string `json:"title" firestore:"title"`
	Desc      string `json:"desc,omitempty" firestore:"desc,omitempty"`
	Completed bool   `json:"completed,omitempty" firestore:"completed,omitempty"`
	Priority  int    `json:"priority,omitempty" firestore:"priority,omitempty"`
	Operator  string `json:"operator,omitempty"`
	AllDay    bool   `json:"allDay,omitempty" firestore:"allDay,omitempty"`
	CSSClass  string `json:"cssClass,omitempty" firestore:"cssClass,omitempty"`
	// Resizable struct {
	// 	gorm.Model
	// 	BeforeStart bool `json:"beforeStart,omitempty" firestore:"beforeStart,omitempty"`
	// 	AfterEnd    bool `json:"afterEnd,omitempty firestore:"afterEnd,omitempty`
	// } `json:"resizable,omitempty" firestore:"resizable,omitempty"`
	Draggable bool      `json:"draggable,omitempty" firestore:"draggable,omitempty"`
	NegiField NegiField `gorm:"foreignKey:id"`
}

func init() {
	db.AutoMigrate(TaskCalEvent{})
}

func (TC *TaskCalEvent) CreateTaskCalEvent() (err error) {

	if err = db.Create(&TC).Error; err != nil {
		return
	}
	return
}

func (TC *TaskCalEvent) UpdateTaskCalEvent() (err error) {

	if err = db.Model(&TC).Updates(&TC).Error; err != nil {
		return
	}

	return
}

func (TC *TaskCalEvent) DeleteTaskCalEvent() (err error) {

	if err = db.Model(&TC).Updates(&TC).Error; err != nil {
		return
	}

	return
}

// func (TC *TaskCalEvent) GetTaskCalEvent() (err error) {

// 	if err = db.Model(&TC).Get()
// }
