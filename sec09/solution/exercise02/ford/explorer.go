package ford

import (
	"errors"
	"fmt"
	"time"

	"github.com/striversity/glft/sec09/solution/exercise02/vehicle"
)

type (
	// Explorer is one of Ford's vehicle that implements vehicle.Vehicle
	Explorer struct {
		vehicleData
	}
	// vehicleData is private to package
	vehicleData struct {
		year  uint16
		seats uint8
	}
)

// NewExplorer returns an initialized ford.Explorer value
func NewExplorer(year uint16) (e *Explorer, err error) {
	// verify validity of year parameter
	currentYear := uint16(time.Now().Year())
	if year < 1990 || year > currentYear {
		return nil, errors.New("invalid year for Ford Explorer")
	}

	e = &Explorer{}
	e.year = year
	e.seats = 5
	return
}

func (e *Explorer) Model() string {
	return "Explorer"
}
func (e *Explorer) Make() string {
	return "Ford"
}
func (e *Explorer) Year() uint16 {
	if e == nil {
		return 0
	}
	return e.year
}
func (e *Explorer) SeatingCap() uint8 {
	if e == nil {
		return 5
	}
	return e.seats
}
func (e *Explorer) Type() int {
	return vehicle.FOUR_DOOR_SUV
}
func (e *Explorer) PowerTrain() string {
	return "gas"
}

func (e *Explorer) String() string {
	var s = fmt.Sprintf(`Model: %v
	Make: %v
	Year: %v
	Type: %v
	Seats: %v,
	Power Train: %v`,
		e.Model(), e.Make(), e.Year(), vehicle.Type(e.Type()), e.SeatingCap(), e.PowerTrain())
	return s
}