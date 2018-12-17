package stats

import "github.com/striversity/glft/shared/sec02/db"

func AvgIncome(getRecord RecordProvider) (income db.Currency) {
	db.ResetCursor() // just incase
	var total db.Currency
	var records uint
	for {
		_, _, _income, err := getRecord()
		if nil != err { // are we are the end?
			break
		}
		records++
		total += _income
	}
	income = total / db.Currency(records)
	return
}
