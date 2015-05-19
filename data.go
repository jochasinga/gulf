/*
 * A (haunted) house should contain...
 *
 * Id			int
 * Address	    string
 * City			string
 * State		string
 * Zipcode		int
 * Built		Time
 * Owner		Owner
 * Status		Status
 * Detail		string
 *
 * 
 * An Owner should contain...
 * Firstname	string
 * Lastname		string
 *
 * A Status should contain...
 * "for rent"	string
 * "for sale"	string
 * "unknown"	string
 *
 */

 package main

/*
 import (
 	"time"
 )
 */

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

 type Owner struct {
 	Firstname	string	`json:"firstname"`
 	Lastname	string	`json:"lastname"`
 }

type Houses []House