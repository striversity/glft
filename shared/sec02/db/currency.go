package db

import (
	"fmt"
)

// Currency is a type to abstract the presentation of dollars
type Currency float64

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", c)
}
