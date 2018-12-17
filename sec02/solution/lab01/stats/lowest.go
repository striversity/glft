package stats

import "github.com/striversity/glft/shared/sec02/db"

func LowestEarner(getRecord RecordProvider) (name, gender string, income db.Currency) {
	db.ResetCursor() // just incase
	for {
		_n, _g, _income, err := getRecord()

		if nil != err { // are we are the end?
			break
		}
		if income > _income || income == 0.0 {
			name = _n
			gender = _g
			income = _income
		}
	}
	return
}
