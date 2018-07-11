package routes

import (
	"errors"
	"holdempoker/models"
)

//AddRoom 룸을 추가 한다.
func AddRoom(data map[string]interface{}) (interface{}, error) {
	var room models.Room
	var returnVal interface{}
	var err error

	if data["buyInMin"] == nil || data["buyInMax"] == nil {
		return nil, errors.New("argument error")
	}

	room = models.Room{}
	room.BuyInMin = int64(data["buyInMin"].(float64))
	room.BuyInMax = int64(data["buyInMax"].(float64))
	room.State = models.RoomStateWait

	session := db.Begin()
	if err := session.Create(&room).Error; err != nil {
		session.Rollback()
		return nil, err
	}
	session.Commit()

	if err := dmap.AddRoom(room); err != nil {
		return nil, err
	}
	return returnVal, err
}

// GetRoom 룸정보를 가져온다.
func GetRoom(data map[string]interface{}) (interface{}, error) {
	var room models.Room
	var returnVal interface{}
	var err error
	if data["roomIndex"] == nil {
		return nil, errors.New("userIndex not found")
	}

	roomIndex := int(data["roomIndex"].(float64))
	room, err = dmap.GetRoom(roomIndex)

	if err != nil {
		return returnVal, err
	}
	returnMap := make(map[string]interface{})
	returnMap["Room"] = room

	gamePlayer := GetGamePlayers(roomIndex)
	returnMap["GamePlayers"] = gamePlayer

	returnVal = returnMap

	return returnVal, nil
}
