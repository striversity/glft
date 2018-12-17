package cart

import (
	"fmt"
	"testing"
)

func TestGetTotal(t *testing.T) {

	for i := 0; i < 10; i++ {
		fmt.Printf("A Cart Total test # %v - got %v\n", i, getTotal())
	}
}
