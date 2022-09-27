package app

import (
	"github.com/satori/go.uuid"
	"time"
)

type NotificationApp struct {
	ID             uuid.UUID `gorm:"PRIMARY_KEY;DEFAULT:uuid_generate_v4()"`
	Name           string    `json:"name" form:"name" query:"name" validate:"required"`
	Alias          string   `gorm:"UNIQUE; NOT NULL" json:"alias" form:"alias" query:"alias" validate:"required"`
	Connect_name   string    `json:"connect_name" form:"connect_name" query:"connect_name" validate:"required"`
	Connect_string string    `json:"connect_string" form:"connect_string" query:"connect_string" validate:"required"`

	created_at time.Time  `json:"created_at" form:"created_at" query:"created_at"`
	updated_at time.Time  `json:"updated_at" form:"updated_at" query:"updated_at"`
	deleted_at *time.Time `json:"deleted_at" form:"deleted_at" query:"deleted_at"`
}
