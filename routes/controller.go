package routes

import (
	database "holdempoker/db"
	"holdempoker/models"

	"github.com/jinzhu/gorm"
)

// Controller is routes Controller
type Controller struct {
	m map[int]interface{}
}

var db *gorm.DB

// Init is 초기화
func (c *Controller) Init() {
	c.m = make(map[int]interface{})
	c.m[1] = Login
	c.m[2] = GetRoom
	db = database.GetDBInstance().GetDB()
}

// Handle 컨트롤러 핸들링
func (c *Controller) Handle(packetNum int, packetData map[string]interface{}) models.PacketData {
	var returnData models.PacketData
	returnData.PacketNum = packetNum

	funcRef, exists := c.m[packetNum]
	if exists {
		returnVal, err := funcRef.(func(map[string]interface{}) (interface{}, error))(packetData)
		if err != nil {
			returnData.Error = err
		}
		returnData.Data = returnVal
	}
	return returnData
}
