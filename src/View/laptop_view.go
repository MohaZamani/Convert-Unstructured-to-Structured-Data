package view

import (
	"github.com/gin-gonic/gin"
	"github.com/mzamani18/rapd_solutions_challenge/entity"
	"github.com/mzamani18/rapd_solutions_challenge/utils"
	responsehandler "github.com/mzamani18/rapd_solutions_challenge/utils/response_handler"
)


func GetAllLaptopsDetails(c *gin.Context){
	all_laptops , err := utils.LoadLaptopDetails()

	if(err!= nil){
		responsehandler.SendUnSucessfulReponse(c, err.Error())
	}

	var laptops_list []*entity.LaptopDetail

	for _, detail := range all_laptops {
		laptops_list = append(laptops_list, detail.LaptopDetail)
	}

	responsehandler.SendSucessFulResponseWithBatchLaptopDetail(c,laptops_list)
}