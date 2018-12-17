package math

import (
	"math"

	"github.com/verrol/glft/sec02/lec12/logger"
)

// Pi is the value of 'math.Pi'
var Pi = math.Pi

func init() {
	logger.Info("my math.init() in my pi.go file")
}
