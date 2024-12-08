package patterns

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	id            uuid.UUID
	userID        uuid.UUID
	from          time.Time
	to            time.Time
	hairDresserID uuid.UUID
}

func CreateBooking(from, to time.Time, userID, hairDresserID uuid.UUID) (*Booking, error) {
	closingTime, _ := time.Parse(time.Kitchen, "17:00pm")

	if from.After(closingTime) {
		return nil, fmt.Errorf("no appointments after %v", closingTime)
	}
	return &Booking{
		hairDresserID: hairDresserID,
		id:            uuid.New(),
		userID:        userID,
		from:          from,
		to:            to,
	}, nil
}
