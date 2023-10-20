package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:'0'"` //0 未完成，1 已完成
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}
