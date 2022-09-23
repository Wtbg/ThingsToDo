package model

import (
	"time"

	// "gopkg.in/guregu/null.v4"
)

type Affair struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"not null"`
	Emergency int `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time
}
