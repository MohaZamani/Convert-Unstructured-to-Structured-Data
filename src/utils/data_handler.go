package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/mzamani18/rapd_solutions_challenge/config"
	"github.com/mzamani18/rapd_solutions_challenge/entity"
)

// Function to load data from the JSON file into a Go map
func LoadLaptopDetails() (map[string]entity.LaptopDetailWithText, error) {
	if _, err := os.Stat(config.Config.LaptopDetailsFileName); os.IsNotExist(err) {
		return make(map[string]entity.LaptopDetailWithText), nil
	}

	file, err := os.Open(config.Config.LaptopDetailsFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var data map[string]entity.LaptopDetailWithText
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Function to write the struct data to a JSON file
func SaveAllLaptopDetails( data map[string]entity.LaptopDetailWithText) error {
	file, err := json.MarshalIndent(data, "", "  ") // Marshal the struct as pretty JSON
	if err != nil {
		return err
	}

	return ioutil.WriteFile(config.Config.LaptopDetailsFileName, file, 0644)
}

// Function to find the next auto-increment ID
func getNextID(data map[string]entity.LaptopDetailWithText) int {
	maxID := 0
	for key := range data {
		id, err := strconv.Atoi(key) 
		if err == nil && id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

// Insert function to add new laptop to existing JSON data with auto-increment key
func InsertLaptopDetail(newLaptop entity.LaptopDetailWithText) error {
	data, err := LoadLaptopDetails()
	if err != nil {
		return err
	}

	// auto-increment mode
	nextID := getNextID(data)

	data[strconv.Itoa(nextID)] = newLaptop

	return SaveAllLaptopDetails(data)
}