package db

// {"id":2,"firstName":"Chere","lastName":"Toyne","gender":"Female","income":"$2.68"},
// person contains fields for a record in our database about an individual
type person struct {
	ID        uint    `json:"id,omitempty"`
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Gender    string  `json:"gender,omitempty"`
	Income    Currency `json:"income,omitempty"`
	Zip       string  `json:"zip,omitempty"`
}
