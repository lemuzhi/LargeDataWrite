package model

import (
	"LargeDataWrite/tools"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	gorm.Model
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate 这里必须用指针，不然id默认给值无法生效，依然还是数据库自增id
func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID, err = tools.GetID()
	return err
}
