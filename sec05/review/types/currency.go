package types

import "fmt"

type Currency float64

func (c Currency) String() string {
	s := fmt.Sprintf("$%.2f", float64(c))
	return s
}
