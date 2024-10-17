package controller

import (
	"github.com/gin-gonic/gin"
	view "github.com/mzamani18/rapd_solutions_challenge/View"
)


func GetAllDocuments(c *gin.Context) {
	view.GetAllLaptopsDetails(c)
}

func SearchOnStructuredData(c *gin.Context) {

	brand := c.Query("brand")
	model := c.Query("model")
	processor := c.Query("processor")
	ramCapacity := c.Query("ram_capacity")
	ramType := c.Query("ram_type")
	storageCapacity := c.Query("storage_capacity")
	batteryStatus := c.Query("battery_status")

	
	view.SearchView(c, brand, model, processor , ramCapacity, ramType, storageCapacity, batteryStatus)
}