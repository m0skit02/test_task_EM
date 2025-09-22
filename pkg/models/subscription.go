package models

import (
	"database/sql/driver"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const humanDateFormat = "02.01.2006"

type HumanDate time.Time

// --- JSON ---
func (d *HumanDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) < 2 {
		return fmt.Errorf("invalid date")
	}
	s = s[1 : len(s)-1]

	t, err := time.Parse(humanDateFormat, s)
	if err != nil {
		return err
	}
	*d = HumanDate(t)
	return nil
}

func (d HumanDate) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	return []byte(fmt.Sprintf("\"%s\"", t.Format(humanDateFormat))), nil
}

// --- SQL (для GORM) ---
func (d HumanDate) Value() (driver.Value, error) {
	return time.Time(d), nil // сохраняем как time.Time
}

func (d *HumanDate) Scan(value interface{}) error {
	if value == nil {
		*d = HumanDate(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = HumanDate(v)
		return nil
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err
		}
		*d = HumanDate(t)
		return nil
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		*d = HumanDate(t)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into HumanDate", value)
	}
}

func (d HumanDate) ToTime() time.Time {
	return time.Time(d)
}

type Subscription struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ServiceName string     `json:"service_name" gorm:"column:service_name;not null"`
	Price       int        `json:"price" gorm:"column:price;not null"`
	UserID      uuid.UUID  `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	StartDate   HumanDate  `json:"start_date" gorm:"column:start_date;not null"`
	EndDate     *HumanDate `json:"end_date,omitempty" gorm:"column:end_date"`
}
