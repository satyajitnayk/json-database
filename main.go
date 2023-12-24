package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger(lumber.INFO)
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}

	// careate db
	opts.Logger.Debug("Creating the database at '%s' ...\n", dir)
	// 0755 - folder access level
	return &driver, os.MkdirAll(dir, 0755)
}

// struct methods
func (d *Driver) Write(collection string, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection. No place to save record!")
	}
	if resource == "" {
		return fmt.Errorf("Missing resource. Unable to save record (no name)!")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	finalPath := filepath.Join(dir, resource+".json")
	tmpPath := finalPath + ".tmp"

	// create collection dir
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	// write to next line
	b = append(b, byte('\n'))

	if err := os.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmpPath, finalPath)
}

func (d *Driver) Read(collection string, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - unable to read!")
	}
	if resource == "" {
		return fmt.Errorf("Missing resource - unable to read record (no name)!")
	}

	record := filepath.Join(d.dir, collection, resource)
	if _, err := stat(record); err != nil {
		return err
	}

	b, err := os.ReadFile(record + ".json")
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("Missing collection - unable to read!")
	}
	dir := filepath.Join(d.dir, collection)
	if _, err := stat(dir); err != nil {
		return nil, err
	}

	files, _ := os.ReadDir(dir)
	var records []string
	for _, file := range files {
		b, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		records = append(records, string(b))
	}
	return records, nil
}

func (d *Driver) Delete(collection string, resource string) error {
	path := filepath.Join(collection, resource)
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, path)

	switch fi, err := stat(dir); {
	case fi == nil, err != nil:
		return fmt.Errorf("unable to find file or directory named %v\n", path)
	case fi.Mode().IsDir():
		return os.RemoveAll(dir)
	case fi.Mode().IsRegular():
		return os.RemoveAll(dir + ".json")
	}
	return nil
}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	m, ok := d.mutexes[collection]
	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}
	return m
}

func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

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
