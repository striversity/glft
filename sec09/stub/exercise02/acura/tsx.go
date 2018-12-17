package acura

import (
	"errors"
	"time"

	"github.com/striversity/glft/sec09/stub/exercise02/vehicle"
)

type (
	// TSX is one of Acura's vehicle that implements vehicle.Vehicle
	TSX struct {
		vehicleData
	}
	// vehicleData is private to package
	vehicleData struct {
		year  uint16
		seats uint8
	}
)

// NewTSX returns an initialized acura.TSX value
func NewTSX(year uint16) (e *TSX, err error) {
	// verify validity of year parameter
	currentYear := uint16(time.Now().Year())
	if year < 2006 || year > currentYear {
		return nil, errors.New("invalid year for Acura TSX")
	}

	e = &TSX{}
	e.year = year
	e.seats = 5
	return
}

func (e *TSX) Model() string {
	return "TSX"
}
func (e *TSX) Make() string {
	return "Acura"
}
func (e *TSX) Year() uint16 {
	if e == nil {
		return 0
	}
	return e.year
}
func (e *TSX) SeatingCap() uint8 {
	if e == nil {
		return 5
	}
	return e.seats
}
func (e *TSX) Type() int {
	return vehicle.FOUR_DOOR_SEDAN
}

// TODO 1 - complete implementation of interface vehicle.Vehice for TSX

// TODO 2 - implement fmt.Stringer for TSX
