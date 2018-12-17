package db

import "testing"

func TestLoad(t *testing.T) {
	testName := "Data successfully loaded"

	t.Run(testName, func(t *testing.T) {
		if err := Load(); err != nil {
			t.Errorf("Load() error = %v, expected nil", err)
		}
	})
}
