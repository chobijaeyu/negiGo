package models

import (
	"gorm.io/gorm"
)

type SeriesTaskSingleTask struct {
	gorm.Model
	Title              string `json:"title,omitempty"`
	Start              int    `json:"start,omitempty"`
	End                int    `json:"end,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	SeriesTaskOptionID uint
}

type SeriesTaskOption struct {
	gorm.Model
	Title    string                 `json:"title,omitempty"`
	Tasklist []SeriesTaskSingleTask `json:"tasklist,omitempty"`
}

func init() {
	db.AutoMigrate(
		SeriesTaskSingleTask{},
		SeriesTaskOption{},
	)
}

func (STO *SeriesTaskOption) CreateSeriesTaskOption() (err error) {
	if err = db.Create(&STO).Error; err != nil {
		return
	}
	return
}

func (STO *SeriesTaskOption) UpdateSeriesTaskOption() (err error) {

	if err = db.Model(&STO).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&STO).Error; err != nil {
		return
	}

	return
}

func (STO *SeriesTaskOption) DeleteSeriesTaskOption() (err error) {

	if err = db.Model(&STO).Delete(&STO).Error; err != nil {
		return
	}

	return
}

// func (STO *SeriesTaskOption) GetSeriesTaskOption() (err error) {

// 	if err = db.Model(&STO).Get()
// }

func (STO *SeriesTaskOption) GetAllSeriesTaskOptions() (cs []SeriesTaskOption, err error) {

	if err = db.Model(&STO).Preload("Tasklist").Find(&cs).Error; err != nil {
		return
	}
	return
}

func (STO *SeriesTaskOption) GetSeriesTaskOptionsByQuery(q, qv, v string) (cvs []SeriesTaskOption, err error) {

	if err = db.Model(&STO).Where(`%v = %v`, q, qv).Find(&cvs).Error; err != nil {
		return
	}
	return
}

func (stst *SeriesTaskSingleTask) DeteleSingleTask() (err error) {
	if err = db.Model(&stst).Delete(&stst).Error; err != nil {
		return
	}
	return
}
