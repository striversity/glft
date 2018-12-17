package main

import (
	"errors"
	"strconv"
	"strings"
)

type (
	// City is a struct with city and country name, and city population
	City struct {
		Name       string
		Country    string
		Population uint
	}
	// Db is a map of country to cities
	Db map[string][]City
)

// CityFromCSV creates a City object from a CSV string.
// Returns an error if 's' doesn't contain the exected fields
func CityFromCSV(s string) (c City, err error) {
	// sample: "Toronto,Canada,5992739"
	const expectedFields = 3
	fields := strings.Split(s, ",")
	if expectedFields > len(fields) {
		return c, errors.New("Incorrect number of fields in 's': " + s)
	}
	c.Name = fields[0]
	c.Country = fields[1]
	p, _ := strconv.ParseUint(fields[2], 10, 64)
	c.Population = uint(p)

	return
}
