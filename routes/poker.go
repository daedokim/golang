package routes

import (
	"errors"
	"holdempoker/models"
	"time"
)

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
		// 데이터맵에 없으면 db에서 조회한다.
		if db.Where("room_index = ?", roomIndex).First(&room).RecordNotFound() == false {
			// db에서 조회한 값이 있다면 데이터맵을 db값으로 동기화한다.
			dmap.AddRoom(room)
		}
	}
	returnMap := make(map[string]interface{})
	returnMap["Room"] = room

	gamePlayer := GetGamePlayers(roomIndex)
	returnMap["GamePlayers"] = gamePlayer

	returnVal = returnMap

	return returnVal, nil
}

// GetGamePlayers 룸내 게임플레이어 정보를 가져온다.
func GetGamePlayers(roomIndex int) []models.GamePlayer {
	var gamePlayers []models.GamePlayer
	gamePlayers, err := dmap.GetGamePlayers(roomIndex)

	if err != nil {
		dbFound := db.Find(&gamePlayers).RecordNotFound()
		if dbFound == true {
			dmap.AddGamePlayers(roomIndex, gamePlayers)
		}
	}
	return gamePlayers
}

//JoinGame 게임테이블에 참여한다.
func JoinGame(data map[string]interface{}) (interface{}, error) {
	var returnVal interface{}

	if data["roomIndex"] == nil || data["userIndex"] == nil || data["chairIndex"] == nil || data["buyInLeft"] == nil {
		return returnVal, errors.New("Argument Error")
	}

	roomIndex, userIndex := int(data["roomIndex"].(float64)), int64(data["userIndex"].(float64))
	chairIndex, buyInLeft := int(data["chairIndex"].(float64)), int64(data["buyInLeft"].(float64))

	if _, err := dmap.GetGamePlayer(roomIndex, userIndex); err != nil {
		newGamePlayer := models.GamePlayer{}
		newGamePlayer.RoomIndex = roomIndex
		newGamePlayer.UserIndex = userIndex
		newGamePlayer.ChairIndex = chairIndex
		newGamePlayer.BuyInLeft = buyInLeft
		newGamePlayer.LastActionDate = time.Now()

		dmap.AddGamePlayer(newGamePlayer)
	} else {
		return returnVal, errors.New("Exist GamePlayer")
	}
	return returnVal, nil
}
