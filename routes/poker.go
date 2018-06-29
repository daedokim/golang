package routes

import (
	"errors"
	"holdempoker/models"
)

// GetRoom 룸정보를 가져온다.
func GetRoom(data map[string]interface{}) (interface{}, error) {
	var room models.Room
	var returnVal interface{}

	if data["roomIndex"] == nil {
		return nil, errors.New("userIndex not found")
	}

	roomIndex := int(data["roomIndex"].(float64))
	db.Where("room_index = ?", roomIndex).First(&room)

	returnMap := make(map[string]interface{})
	returnMap["Room"] = room
	returnMap["GamePlayers"] = GetGamePlayers(roomIndex)

	returnVal = returnMap

	return returnVal, nil
}

// GetGamePlayers 룸내 게임플레이어 정보를 가져온다.
func GetGamePlayers(roomIndex int) []models.GamePlayer {
	var gamePlayers []models.GamePlayer
	if err := db.Find(&gamePlayers).Error; err != nil {
		return nil
	}
	return gamePlayers
}
