package responsehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mzamani18/rapd_solutions_challenge/entity"
	openaiservice "github.com/mzamani18/rapd_solutions_challenge/open_ai_services"
	"github.com/sashabaranov/go-openai"
)

func SendBadRequest(c *gin.Context, message string)  {
	c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
	})
}

func SendSucessFulResponseWithLaptopDetail(c *gin.Context,laptop_detail *entity.LaptopDetail){
	c.JSON(http.StatusOK, gin.H{
		"brand:" : laptop_detail.Brand,
		"model": laptop_detail.Model,
		"processor": laptop_detail.Processor,
		"ram_capacity": laptop_detail.RamCapacity,
		"ram_type" : laptop_detail.RamType,
		"storage_capacity": laptop_detail.StorageCapacity,
		"battery_status" : laptop_detail.BatteryStatus,
	})
}


func SendSucessFulResponseWithBatchLaptopDetail(c *gin.Context, laptops_details []*entity.LaptopDetail){
	var laptops []gin.H

	for _, laptop_detail := range laptops_details {
		laptops = append(laptops, gin.H{
			"brand":           laptop_detail.Brand,
			"model":           laptop_detail.Model,
			"processor":       laptop_detail.Processor,
			"ram_capacity":    laptop_detail.RamCapacity,
			"ram_type":        laptop_detail.RamType,
			"storage_capacity": laptop_detail.StorageCapacity,
			"battery_status":  laptop_detail.BatteryStatus,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"laptops": laptops,
	})
}

func SendUnSucessfulReponse(c *gin.Context, message string){
	c.JSON(http.StatusNoContent, gin.H{
		"message": message,
	})
}

func SendUnSucessfulReponseFromOpenAI(c *gin.Context, err openai.APIError){
	openaiservice.OpenAIErrorHandler(err, c)
}