package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	utils "github.com/mzamani18/rapd_solutions_challenge/utils"
	responsehandler "github.com/mzamani18/rapd_solutions_challenge/utils/response_handler"
	"github.com/sashabaranov/go-openai"
)

func ConvertDataToStructuredData(c *gin.Context){
	type unstructured_data struct {
		Text string `json:"text" binding:"required"`
	}

	var req_body unstructured_data

	if c.Bind(&req_body) != nil {
		responsehandler.SendBadRequest(c, "Bad Request!")
		return
	}

	laptop_detail, err := utils.ConvertTextToStructuredData(strings.ToLower(req_body.Text))

	if(err!= nil){
		if apiErr, ok := err.(*openai.APIError); ok {
			responsehandler.SendUnSucessfulReponseFromOpenAI(c, *apiErr)
		}else{
			responsehandler.SendBadRequest(c, err.Error())
		}
		return
	}

	responsehandler.SendSucessFulResponseWithLaptopDetail(c, laptop_detail)
}


func ConvertBatchDataToStructuredData(c *gin.Context){

	type unstructured_data struct {
		Texts []string `json:"texts" binding:"required"`
	}

	var req_body unstructured_data

	if c.Bind(&req_body) != nil {
		responsehandler.SendBadRequest(c, "Bad Request!")
		return
	}

	for ind, text:= range req_body.Texts{
		req_body.Texts[ind] = strings.ToLower(text)
	}

	laptop_detail, err := utils.ConvertBatchTextToStructuredData(req_body.Texts)

	if(err!= nil){
		if apiErr, ok := err.(*openai.APIError); ok {
			responsehandler.SendUnSucessfulReponseFromOpenAI(c, *apiErr)
		}else{
			responsehandler.SendBadRequest(c, err.Error())
		}
		return
	}

	responsehandler.SendSucessFulResponseWithBatchLaptopDetail(c, laptop_detail)
}