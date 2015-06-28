package models

type Status int

// Enum for status
const (
	Unknown Status = iota
	Sale
	Rent 	
	Airbnb 	
)

var statuses = [...]string{
	"Unknown",
	"For Sale",
	"For Rent",
	"On Airbnb",
}

func (s Status) String() string {
	return statuses[s]
}

type House struct {
	Id		int			`json:"id"`
	Address	string		`json:"address"`
	City	string		`json:"city"`
	State	string		`json:"state"`
	Zipcode string		`json:"zipcode"`
	Built	int			`json:"built"`
	Owner	Owner		`json:"owner"`
	Status	Status		`json:"status"`
	Detail 	string		`json:"detail"`
}