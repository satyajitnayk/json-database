package main

import (
	"encoding/json"
	"fmt"

	"github.com/satyajitnayk/json-database/database"
	"github.com/satyajitnayk/json-database/models"
)

func main() {
	dir := "./"
	db, err := database.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []models.User{
		{"John", "23", "1234567890", "MI Tech", models.Address{"banglaore", "karnataka", "india", "410233"}},
		{"Alice", "25", "9876543210", "Tech Innovators", models.Address{"Mumbai", "Maharashtra", "India", "400001"}},
		{"Bob", "30", "8765432109", "Data Wizards", models.Address{"Delhi", "Delhi", "India", "110001"}},
		{"Eva", "28", "7654321098", "Code Masters", models.Address{"Chennai", "Tamil Nadu", "India", "600001"}},
		{"David", "35", "6543210987", "Byte Builders", models.Address{"Hyderabad", "Telangana", "India", "500001"}},
		{"Sophie", "22", "5432109876", "Byte Crafters", models.Address{"Pune", "Maharashtra", "India", "411001"}},
		{"Michael", "27", "4321098765", "Logic Loom", models.Address{"Kolkata", "West Bengal", "India", "700001"}},
		{"Emma", "29", "3210987654", "Quantum Solutions", models.Address{"Ahmedabad", "Gujarat", "India", "380001"}},
		{"Daniel", "32", "2109876543", "Infinite Innovations", models.Address{"Jaipur", "Rajasthan", "India", "302001"}},
		{"Grace", "26", "1098765432", "Data Dynamics", models.Address{"Lucknow", "Uttar Pradesh", "India", "226001"}},
		{"Ryan", "31", "9876543210", "Tech Titans", models.Address{"Bengaluru", "Karnataka", "India", "560001"}},
		{"Olivia", "24", "8765432109", "Code Crafters", models.Address{"Chandigarh", "Chandigarh", "India", "160001"}},
		{"Matthew", "33", "7654321098", "Innovate Minds", models.Address{"Coimbatore", "Tamil Nadu", "India", "641001"}},
		{"Lily", "28", "6543210987", "Data Dreamers", models.Address{"Nagpur", "Maharashtra", "India", "440001"}},
		{"Andrew", "30", "5432109876", "Byte Innovations", models.Address{"Indore", "Madhya Pradesh", "India", "452001"}},
		{"Aiden", "25", "4321098765", "Quantum Creators", models.Address{"Bhopal", "Madhya Pradesh", "India", "462001"}},
	}

	for _, value := range employees {
		// use the username to create file for the collection
		db.Write("users", value.Name, models.User{
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

	allusers := []models.User{}

	for _, f := range records {
		employeeFound := models.User{}
		// data is in json so unmarshal it for golang to understand
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}

	fmt.Println(allusers)

	if err := db.Delete("users", "john"); err != nil {
		fmt.Println("Error", err)
	}

	if err := db.Delete("users", ""); err != nil {
		fmt.Println("Error", err)
	}
}
