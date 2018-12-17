package stats

import "github.com/striversity/glft/shared/sec02/db"

type RecordProvider func() (name, gender string, income db.Currency, err error)
