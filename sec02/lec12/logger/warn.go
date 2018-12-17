package logger

import (
	"fmt"
)

// Warn is a simple log function
func Warn(s string) {
	fmt.Println("[Warn]", s)
}

func init() {
	fmt.Println("logger.init() - 1 from warn.go")
}
