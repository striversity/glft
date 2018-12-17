package logger

import (
	"fmt"
)

// Info is a simple log function
func Info(s string) {
	fmt.Println("[INFO]", s)
}

func init() {
	fmt.Println("logger.init() - 1 from info.go")
}
