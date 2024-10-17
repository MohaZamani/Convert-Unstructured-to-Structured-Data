package view

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mzamani18/rapd_solutions_challenge/entity"
	"github.com/mzamani18/rapd_solutions_challenge/utils"
	responsehandler "github.com/mzamani18/rapd_solutions_challenge/utils/response_handler"
)


func SearchView(c *gin.Context, brand string, model string, processor string, ramCapacity string, ramType string, storageCapacity string,batteryStatus string){

	var results []*entity.LaptopDetail

	all_laptops , err := utils.LoadLaptopDetails()

	if(err!= nil){
		responsehandler.SendUnSucessfulReponse(c, err.Error())
	}
	
	// Iterate through all laptop details and filter based on query params
	for _, laptop := range all_laptops {
		if (brand == "" || strings.Contains(strings.ToLower(laptop.LaptopDetail.Brand), brand)) &&
			(model == "" || strings.Contains(strings.ToLower(laptop.LaptopDetail.Model), model)) &&
			(processor == "" || strings.Contains(strings.ToLower(laptop.LaptopDetail.Processor), processor)) &&
			(ramCapacity == "" || strings.Contains(strings.ToLower(laptop.LaptopDetail.RamCapacity), ramCapacity)) &&
			(ramType == "" || strings.Contains(strings.ToLower(laptop.LaptopDetail.RamType), ramType)) &&
			(storageCapacity == "" || strings.Contains(strings.ToLower(laptop.LaptopDetail.StorageCapacity), storageCapacity)) &&
			(batteryStatus == "" || strings.Contains(strings.ToLower(laptop.LaptopDetail.BatteryStatus), batteryStatus)) {

			results = append(results, laptop.LaptopDetail)

		}
	}

	responsehandler.SendSucessFulResponseWithBatchLaptopDetail(c, results)
}