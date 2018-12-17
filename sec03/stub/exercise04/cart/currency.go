package cart

import (
	"fmt"
)

// Currency is an abstraction for money type
type Currency float64

func (c Currency) String() string {
	s := fmt.Sprintf("$%.2f", float64(c))
	return s
}
