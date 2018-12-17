package vehicle

type (

	// Vehicle interface is implemented by any value that has the following methods
	Vehicle interface {
		// Model returns the model name of the type of vehicle, for example Explorer, Model3, etc
		Model() string
		// Make returns the make of the vehicle, for example Ford, Tesla, etc.
		Make() string
		// Year returns the model year between 1910 and current year, for example 2018
		Year() uint16
		// SeatingCap is the nubmer of passengers seating avaiable
		SeatingCap() uint8
		// Type returns the type of vehicle as defined in pacakge vehicle
		Type() int
		// PowerTrain returns either electric, hybrid, gas, diesel
		PowerTrain() string
	}
)

// Type defines possible car types
const (
	FOUR_DOOR_SEDAN  = iota
	FOUR_DOOR_SUV    = iota
	TWO_DOOR_PICKUP  = iota
	FOUR_DOOR_PICKUP = iota
)

// Type returns a string for the type code
func Type(t int) string {
	switch t {
	case FOUR_DOOR_SEDAN:
		return "4 Door Sedan"
	case FOUR_DOOR_SUV:
		return "4 Door SUV"
	case TWO_DOOR_PICKUP:
		return "2 Door Pickup"
	case FOUR_DOOR_PICKUP:
		return "4 Door Pickup"
	default:
		return "illegal vehicle type"
	}
}
