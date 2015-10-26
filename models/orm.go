package models

import (
	"github.com/jochasinga/go-routing/db"
)

var currentHouseId int

/**** House's method ****/

func (h *House) Save() {
	// Save to data slice
	currentHouseId += 1
	h.Id = currentHouseId
	houses = append(db.houses, h)
}

/**** House's helper functions ****/

func FindAll() Houses {
	return db.houses
}

func Create(h *House) {
	h.Save()
}

func House(args ...interface{}) (*House) {
	
	house := new(House)
	
	for index, arg := range args {
		switch v := arg.(type) {
		case string && index == 0:
			house.Address = v
		case string && index == 1:
			house.City = v
		case string && index == 2:
			house.State = v
		case string && index == 3:
			house.Zipcode = v
		case int && index == 4:
			house.Built = v
		case Owner && index == 5:
			house.Owner = v
		case Status && index == 6:
			house.Status = v
		case string && index == 7:
			house.Detail = v
		}
	}
	return house
}

func (h *House) Save() {
	// Save to data slice
	db.houses = append(db.houses, h)
}

func FindById(id int) House {
	var h House 

	for _, house := range db.houses {
		if house.Id == id {
			h = house 
		}
	}
	return h
}

func FindByCity(city string) Houses {
	var hs Houses 

	for _, house := range db.houses {
		if house.City == city {
			hs = append(hs, house)
		}
	}
	return hs 
}

func FindByState(state string) Houses {
	var hs Houses

	for _, house := range db.houses {
		if house.State == state {
			hs = append(hs, house)
		}
	}
	return hs
}