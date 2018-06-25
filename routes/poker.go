package routes

import (
	"fmt"
	. "holdempoker/models"
)

// Login is login
func Login(packetData interface{}) bool {

	returnVal := false

	return returnVal
}

// GetRoom is Get RoomInfo
func GetRoom(packetData interface{}) interface{} {
	room := Room{}
	fmt.Printf("recv:%#v\n", packetData)

	return room
}
