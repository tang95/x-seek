package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Direction string

const (
	Asc  Direction = "ASC"
	Desc Direction = "DESC"
)

type PageQuery struct {
	Page int32 `json:"page"`
	Size int32 `json:"size"`
}

func (pagination *PageQuery) GetOffset() int {
	return int((pagination.Page - 1) * pagination.Size)
}

func (pagination *PageQuery) GetLimit() int {
	return int(pagination.Size)
}

type SortQuery struct {
	Field     string    `json:"field"`
	Direction Direction `json:"direction"`
}

type BaseModel struct {
	ID        string    `gorm:"primaryKey; type:varchar(36); not null" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime; not null; comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime; not null; comment:更新时间" json:"updated_at"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}
