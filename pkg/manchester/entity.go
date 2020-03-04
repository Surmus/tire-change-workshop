package manchester

import (
	"time"
)

var zeroTireChangeTimeEntity = &tireChangeTimeEntity{}

type tireChangeTimeEntity struct {
	ID uint `gorm:"primary_key"`

	Time time.Time

	Available bool

	BookedByContact string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func newTireChangeTimeEntity(changeTime time.Time, available bool) *tireChangeTimeEntity {
	return &tireChangeTimeEntity{
		Time:      changeTime,
		Available: available,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (e *tireChangeTimeEntity) makeBooking(contactInformation string) {
	if e == zeroTireChangeTimeEntity || !e.Available {
		panic(newUnAvailableBookingError())
	}

	e.Available = false
	e.UpdatedAt = time.Now()
	e.BookedByContact = contactInformation
}

func (e tireChangeTimeEntity) TableName() string {
	return "tire_change_time"
}
