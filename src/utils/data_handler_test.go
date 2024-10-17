package utils

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/mzamani18/rapd_solutions_challenge/config"
	"github.com/mzamani18/rapd_solutions_challenge/entity"
)

var testDataFile = "test_laptop_details.json"

// Setup function to prepare the test environment
func setup() {
	config.Config.LaptopDetailsFileName = testDataFile
	os.Remove(testDataFile) // Clean up any existing test data
}

// Teardown function to clean up after tests
func teardown() {
	os.Remove(testDataFile)
}

func TestLoadLaptopDetails_FileNotExists(t *testing.T) {
	setup()
	defer teardown()

	data, err := LoadLaptopDetails()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(data) != 0 {
		t.Error("Expected empty map when file does not exist")
	}
}

func TestLoadLaptopDetails_Success(t *testing.T) {
	setup()
	defer teardown()

	// Create sample data to save
	sampleData := map[string]entity.LaptopDetailWithText{
		"1": {
			Text:          "Laptop: Dell Inspiron",
			LaptopDetail: &entity.LaptopDetail{
				Brand:          "Dell",
				Model:          "Inspiron",
				Processor:      "i7",
				RamCapacity:    "16GB",
				RamType:        "DDR4",
				StorageCapacity: "512GB SSD",
				BatteryStatus:  "Yes",
			},
		},
	}

	// Save the sample data
	SaveAllLaptopDetails(sampleData)

	// Now load the data
	loadedData, err := LoadLaptopDetails()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(loadedData) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(loadedData))
	}
}

func TestSaveAllLaptopDetails(t *testing.T) {
	setup()
	defer teardown()

	data := map[string]entity.LaptopDetailWithText{
		"1": {
			Text:          "Laptop: HP Spectre",
			LaptopDetail: &entity.LaptopDetail{
				Brand:          "HP",
				Model:          "Spectre",
				Processor:      "i5",
				RamCapacity:    "8GB",
				RamType:        "LPDDR4",
				StorageCapacity: "256GB SSD",
				BatteryStatus:  "No",
			},
		},
	}

	// Save the data
	err := SaveAllLaptopDetails(data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the data was saved correctly
	fileData, err := os.ReadFile(testDataFile)
	if err != nil {
		t.Fatalf("Expected no error reading file, got %v", err)
	}

	var loadedData map[string]entity.LaptopDetailWithText
	if err := json.Unmarshal(fileData, &loadedData); err != nil {
		t.Fatalf("Expected no error unmarshalling file data, got %v", err)
	}

	if len(loadedData) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(loadedData))
	}
}

func TestGetNextID(t *testing.T) {
	data := map[string]entity.LaptopDetailWithText{
		"1": {},
		"2": {},
		"5": {},
	}

	nextID := getNextID(data)
	if nextID != 6 {
		t.Errorf("Expected next ID to be 6, got %d", nextID)
	}

	// Test with no existing IDs
	data = make(map[string]entity.LaptopDetailWithText)
	nextID = getNextID(data)
	if nextID != 1 {
		t.Errorf("Expected next ID to be 1, got %d", nextID)
	}
}

func TestInsertLaptopDetail(t *testing.T) {
	setup()
	defer teardown()

	newLaptop := entity.LaptopDetailWithText{
		Text: "Laptop: Apple MacBook Pro",
		LaptopDetail: &entity.LaptopDetail{
			Brand:          "Apple",
			Model:          "MacBook Pro",
			Processor:      "M1",
			RamCapacity:    "16GB",
			RamType:        "LPDDR4X",
			StorageCapacity: "512GB SSD",
			BatteryStatus:  "Yes",
		},
	}

	err := InsertLaptopDetail(newLaptop)
	if err != nil {
		t.Fatalf("Expected no error when inserting laptop detail, got %v", err)
	}

	// Load the data back to verify
	loadedData, err := LoadLaptopDetails()
	if err != nil {
		t.Fatalf("Expected no error when loading laptop details, got %v", err)
	}

	if len(loadedData) != 1 {
		t.Errorf("Expected 1 entry after insert, got %d", len(loadedData))
	}
}
