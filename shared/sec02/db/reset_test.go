package db

import (
	"testing"
)

func TestResetCursor(t *testing.T) {
	testName := "Test resetting cursor"
	firstTime := true
	count := make(map[bool]int)
	t.Run(testName, func(t *testing.T) {
		err := Load()

		if nil != err {
			t.Errorf("loading data failed - error %v", err)
			return
		}

		for {
			_, _, _, _, _, err := GetNextRecord()
			if err == nil {
				count[firstTime]++
				continue
			}

			if firstTime {
				firstTime = false
				ResetCursor()
			} else {
				break
			}
		}

		if count[true] != count[false] {
			t.Errorf("Reset failed - error mismatch count, expected %v == %v", count[true], count[false])
		}
	})
}
