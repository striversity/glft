package honda

import (
	"errors"
	"time"

	"github.com/striversity/glft/sec09/stub/exercise02/vehicle"
)

type (
	// Pilot is one of Honda's vehicle that implements vehicle.Vehicle
	Pilot struct {
		vehicleData
	}
	// vehicleData is private to package
	vehicleData struct {
		year  uint16
		seats uint8
	}
)

// NewPilot returns an initialized honda.Pilot value
func NewPilot(year uint16) (e *Pilot, err error) {
	// verify validity of year parameter
	currentYear := uint16(time.Now().Year())
	if year < 2006 || year > currentYear {
		return nil, errors.New("invalid year for Honda Pilot")
	}

	e = &Pilot{}
	e.year = year
	e.seats = 8
	return
}

func (e *Pilot) Model() string {
	return "Pilot"
}
func (e *Pilot) Make() string {
	return "Honda"
}
func (e *Pilot) Year() uint16 {
	if e == nil {
		return 0
	}
	return e.year
}
func (e *Pilot) SeatingCap() uint8 {
	if e == nil {
		return 8
	}
	return e.seats
}
func (e *Pilot) Type() int {
	return vehicle.FOUR_DOOR_SUV
}

// TODO 1 - complete implementation of interface vehicle.Vehice for Pilot

// TODO 2 - implement fmt.Stringer for Pilot
