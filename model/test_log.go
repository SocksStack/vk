package model

import "gorm.io/gorm"

type TestLog struct {
	gorm.Model
}

func (m TestLog) TableName() string {
	return "test_log"
}