package routes

import (
	"holdempoker/models"
)

// GetRoom is Get RoomInfo
func GetRoom(data map[string]interface{}) (interface{}, error) {
	room := models.Room{}
	//fmt.Printf("recv:%#v\n", packetData)

	return room, nil
}
