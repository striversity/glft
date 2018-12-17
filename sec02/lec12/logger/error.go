package logger

import (
	"fmt"
)

func init() {
	fmt.Println("logger.init() - 1 from error.go")
}
// Error is a simple log function
func Error(s string) {
	fmt.Println("[ERROR]", s)
}
