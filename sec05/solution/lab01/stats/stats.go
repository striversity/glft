package stats

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/sec05/solution/lab01/types"
)

type (
	myStats struct {
		total     int
		minSalary types.Currency
		maxSalary types.Currency
		avgSalary types.Currency
	}
)

var (
	maleStats, femaleStats myStats
)

// Print takes a Db and computes a few stats:
// total, min, avg, max for females and males in the data set.
func Print(db types.Db) {
	if nil == db {
		log.Info("Empty database, no stats to compute")
		return
	}
	maleStats = genderStatInit(db['M'])
	femaleStats = genderStatInit(db['F'])
	femaleStats.print("Female")
	maleStats.print("Male")
}
func (s myStats) print(h string) {
	fmt.Printf("%v Stats:\n", h)
	fmt.Printf("\tCount: %v\n", s.total)
	fmt.Printf("\tMin Salary: %v\n", s.minSalary)
	fmt.Printf("\tMax Salary: %v\n", s.maxSalary)
	fmt.Printf("\tAvg Salary: %v\n", s.avgSalary)
}
func genderStatInit(p types.People) (g myStats) {
	g.total = len(p)
	g.maxSalary, g.minSalary, g.avgSalary = getMaxSalary(p)
	return
}

func getMaxSalary(p types.People) (max, min, avg types.Currency) {
	var s types.Currency
	for _, v := range p {
		s += v.Salary
		if max < v.Salary {
			max = v.Salary
		}
		if min > v.Salary || min == 0 {
			min = v.Salary
		}
	}
	avg = s / types.Currency(len(p))
	return
}
