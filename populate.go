package main 

// Populate database with mock data
func init() {

	mockhouses := [...]House{

		House{
			Address: "5873 Delaware Avenue",
			City: "Lady Lake",
			State: "FL",
			Zipcode: "32159",
			Built: time.Now().Year() - 60,
			Owner: Owner{ "John", "Jones", },
			Status: Sale,
			Detail: "Just another haunted house.",
		},

		House{
			Address: "8352 Parker Street",
			City: "Salisbury",
			State: "MD", 
			Zipcode: "21801",
			Built: time.Now().Year() - 100,
			Owner: Owner{ "Peter", "Miller", },
			Status: Airbnb,
			Detail: "A house with a ghost of soldier coming home.",
		},

		House{
			Address: "2083 Country Club Drive",
			City: "Leesburg", 
			State: "VA", 
			Zipcode: "20175", 
			Built: time.Now().Year() - 40,
			Owner: Owner{ "Mary", "Lyndel", },
			Detail: "People spot a woman in black with candle at night.",
		},

		House{
			Address: "9347 Surrey Lane",
			City: "Apple Valley", 
			State: "CA", 
			Zipcode: "92307", 
			Built: time.Now().Year() - 20,
			Owner: Owner{ "Michael", "Powell", },
			Detail: "An unrest crime scene of a massacre.",
		},
	}

	for _, house := range mockhouses {
		Create(house)
	}
}