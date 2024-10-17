package utils

import (
	"github.com/mzamani18/rapd_solutions_challenge/entity"
	openaiservice "github.com/mzamani18/rapd_solutions_challenge/open_ai_services"
)


func ConvertTextToStructuredData(text string) (*entity.LaptopDetail, error){
	// here if text exist, should not send request to open ai
	if(trie.Exist(text)){
		return trie.GetLaptopDetail(text), nil
	}

	// get laptop details from open ai
	laptop_detail , err := openaiservice.GetLaptopDetail(text)

	if(err != nil){
		return nil, err
	}

	// save laptop details data to a json file
	err = InsertLaptopDetail(entity.LaptopDetailWithText{LaptopDetail: laptop_detail, Text: text})
	if(err != nil){
		return nil, err
	}

	// insert data to trie
	trie.Insert(text)
	trie.SetLapTopDetail(text, laptop_detail)

	return laptop_detail, nil
}


func ConvertBatchTextToStructuredData(texts []string)([]*entity.LaptopDetail, error){
	var all_laptops_details []*entity.LaptopDetail

	// Check if the input already exists in the trie to avoid redundant OpenAI requests
	for i := 0; i < len(texts); i++ {
		if(trie.Exist(texts[i])){
			all_laptops_details = append(all_laptops_details, trie.GetLaptopDetail(texts[i]))
			texts = append(texts[:i], texts[i+1:]...)
		}
	}

	for _, text:= range texts{
		// get laptop details from open ai
		laptop_detail, err := openaiservice.GetLaptopDetail(text)

		if(err != nil){
			return nil, err
		}

		all_laptops_details = append(all_laptops_details, laptop_detail)

		err = InsertLaptopDetail(entity.LaptopDetailWithText{LaptopDetail: laptop_detail, Text: text})

		if(err != nil){
			return nil, err
		}

		trie.Insert(text)
		trie.SetLapTopDetail(text, laptop_detail)
	}

	return all_laptops_details, nil
}

