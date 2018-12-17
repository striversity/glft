package tesla

import (
	"errors"
	"fmt"
	"time"

	"github.com/striversity/glft/sec09/solution/exercise02/vehicle"
)

type (
	// Model3 is one of Tesla's vehicle that implements vehicle.Vehicle
	Model3 struct {
		vehicleData
	}
	// vehicleData is private to package
	vehicleData struct {
		year  uint16
		seats uint8
	}
)

// NewModel3 returns an initialized tesla.Model3 value
func NewModel3(year uint16) (e *Model3, err error) {
	// verify validity of year parameter
	currentYear := uint16(time.Now().Year())
	if year < 2018 || year > currentYear {
		return nil, errors.New("invalid year for Tesla Model 3")
	}

	e = &Model3{}
	e.year = year
	e.seats = 5
	return
}

func (e *Model3) Model() string {
	return "Model 3"
}
func (e *Model3) Make() string {
	return "Tesla"
}
func (e *Model3) Year() uint16 {
	if e == nil {
		return 0
	}
	return e.year
}
func (e *Model3) SeatingCap() uint8 {
	if e == nil {
		return 5
	}
	return e.seats
}
func (e *Model3) Type() int {
	return vehicle.FOUR_DOOR_SEDAN
}
func (e *Model3) PowerTrain() string {
	return "electric"
}

func (e *Model3) String() string {
	var s = fmt.Sprintf(`Model: %v
	Make: %v
	Year: %v
	Type: %v
	Seats: %v,
	Power Train: %v`,
		e.Model(), e.Make(), e.Year(), vehicle.Type(e.Type()), e.SeatingCap(), e.PowerTrain())
	return s
}
