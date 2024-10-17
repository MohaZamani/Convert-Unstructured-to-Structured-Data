package entity


type LaptopDetail struct {
	Brand string
	Model string 
	Processor  string
	RamCapacity string
	RamType string 
   	StorageCapacity string
   	BatteryStatus string 
}

type LaptopDetailWithText struct {
	LaptopDetail *LaptopDetail
	Text string
}
