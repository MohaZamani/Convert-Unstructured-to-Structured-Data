package utils

import (
	"testing"

	"github.com/mzamani18/rapd_solutions_challenge/entity"
)

func TestInsertAndExist(t *testing.T) {
	InitializeTrie() // Make sure to initialize your trie

	testStrings := []string{
		"Laptop: Dell Inspiron",
		"Laptop: HP Spectre",
		"Laptop: Apple MacBook Pro",
	}

	// Insert test strings into the Trie
	for _, str := range testStrings {
		trie.Insert(str)
	}

	// Check if the strings exist in the Trie
	for _, str := range testStrings {
		if !trie.Exist(str) {
			t.Errorf("Expected %s to exist in the Trie, but it doesn't", str)
		}
	}

	// Check a string that was not added
	if trie.Exist("Laptop: Lenovo ThinkPad") {
		t.Error("Expected Laptop: Lenovo ThinkPad to NOT exist in the Trie, but it does")
	}
}

func TestSetAndGetLaptopDetail(t *testing.T) {
	InitializeTrie() // Make sure to initialize your trie

	laptopDetail := &entity.LaptopDetail{
		Brand:          "Dell",
		Model:          "Inspiron",
		Processor:      "i7",
		RamCapacity:    "16GB",
		RamType:        "DDR4",
		StorageCapacity: "512GB SSD",
		BatteryStatus:  "Yes",
	}

	// The text to insert into the Trie
	text := "Laptop: Dell Inspiron"

	// Insert the text and set the laptop detail
	trie.Insert(text)
	trie.SetLapTopDetail(text, laptopDetail)

	// Retrieve the laptop detail from the Trie
	retrievedDetail := trie.GetLaptopDetail(text)

	if retrievedDetail == nil {
		t.Fatal("Expected to retrieve laptop detail, but got nil")
	}

	if retrievedDetail.Brand != laptopDetail.Brand {
		t.Errorf("Expected brand %s, but got %s", laptopDetail.Brand, retrievedDetail.Brand)
	}
	if retrievedDetail.Model != laptopDetail.Model {
		t.Errorf("Expected model %s, but got %s", laptopDetail.Model, retrievedDetail.Model)
	}
}

func TestGetLaptopDetailNonExistent(t *testing.T) {
	InitializeTrie() // Make sure to initialize your trie

	text := "Laptop: Non-existent"

	// Attempt to get a laptop detail that does not exist
	detail := trie.GetLaptopDetail(text)

	if detail != nil {
		t.Error("Expected nil for non-existent laptop detail, but got something")
	}
}
