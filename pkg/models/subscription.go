package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ServiceName string     `json:"service_name" gorm:"column:service_name;not null"`
	Price       int        `json:"price" gorm:"column:price;not null"`
	UserID      uuid.UUID  `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	StartDate   time.Time  `json:"start_date" gorm:"column:start_date;not null"`
	EndDate     *time.Time `json:"end_date,omitempty" gorm:"column:end_date"`
}
