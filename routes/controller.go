package routes

import (
	database "holdempoker/db"
	"holdempoker/maps"
	"holdempoker/models"

	"github.com/jinzhu/gorm"
)

// Controller is routes Controller
type Controller struct {
	m map[int]interface{}
}

var db *gorm.DB
var dmap *maps.DataMap

// Init is 초기화
func (c *Controller) Init() {
	c.m = make(map[int]interface{})
	db = database.GetInstance().GetDB()
	dmap = maps.GetInstance()

	c.m[1] = Login
	c.m[2] = GetRoom
	c.m[3] = Sit
	c.m[4] = AddRoom
	c.m[5] = StandUp
	c.m[6] = SetPlayerBet
}

// Handle 컨트롤러 핸들링
func (c *Controller) Handle(packetNum int, packetData map[string]interface{}) models.PacketData {
	var returnData models.PacketData
	returnData.PacketNum = packetNum

	funcRef, exists := c.m[packetNum]
	if exists {
		returnVal, err := funcRef.(func(map[string]interface{}) (interface{}, error))(packetData)
		if err != nil {
			returnData.Error = err.Error()
		}
		returnData.Data = returnVal
	}
	return returnData
}
