package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"
	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "23", "1234567890", "MI Tech", Address{"banglaore", "karnataka", "india", "410233"}},
		{"Alice", "25", "9876543210", "Tech Innovators", Address{"Mumbai", "Maharashtra", "India", "400001"}},
		{"Bob", "30", "8765432109", "Data Wizards", Address{"Delhi", "Delhi", "India", "110001"}},
		{"Eva", "28", "7654321098", "Code Masters", Address{"Chennai", "Tamil Nadu", "India", "600001"}},
		{"David", "35", "6543210987", "Byte Builders", Address{"Hyderabad", "Telangana", "India", "500001"}},
		{"Sophie", "22", "5432109876", "Byte Crafters", Address{"Pune", "Maharashtra", "India", "411001"}},
		{"Michael", "27", "4321098765", "Logic Loom", Address{"Kolkata", "West Bengal", "India", "700001"}},
		{"Emma", "29", "3210987654", "Quantum Solutions", Address{"Ahmedabad", "Gujarat", "India", "380001"}},
		{"Daniel", "32", "2109876543", "Infinite Innovations", Address{"Jaipur", "Rajasthan", "India", "302001"}},
		{"Grace", "26", "1098765432", "Data Dynamics", Address{"Lucknow", "Uttar Pradesh", "India", "226001"}},
		{"Ryan", "31", "9876543210", "Tech Titans", Address{"Bengaluru", "Karnataka", "India", "560001"}},
		{"Olivia", "24", "8765432109", "Code Crafters", Address{"Chandigarh", "Chandigarh", "India", "160001"}},
		{"Matthew", "33", "7654321098", "Innovate Minds", Address{"Coimbatore", "Tamil Nadu", "India", "641001"}},
		{"Lily", "28", "6543210987", "Data Dreamers", Address{"Nagpur", "Maharashtra", "India", "440001"}},
		{"Andrew", "30", "5432109876", "Byte Innovations", Address{"Indore", "Madhya Pradesh", "India", "452001"}},
		{"Aiden", "25", "4321098765", "Quantum Creators", Address{"Bhopal", "Madhya Pradesh", "India", "462001"}},
	}

	for _, value := range employees {
		// use the username to create file for the collection
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		// data is in json so unmarshal it for golang to understand
		if err != json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}
}
