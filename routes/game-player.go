package routes

import (
	"errors"
	"holdempoker/models"
	"time"
)

// GetGamePlayers 룸내 게임플레이어 정보를 가져온다.
func GetGamePlayers(roomIndex int) []models.GamePlayer {
	var gamePlayers []models.GamePlayer
	gamePlayers, err := dmap.GetGamePlayers(roomIndex)

	if err != nil {
		notFound := db.Find(&gamePlayers).RecordNotFound()
		if notFound == true {
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

	var user models.User
	userNotFound := db.Where(&models.User{UserIndex: userIndex}).First(&user).RecordNotFound()

	if _, err := dmap.GetGamePlayer(roomIndex, userIndex); err != nil {
		newGamePlayer := models.GamePlayer{}
		newGamePlayer.RoomIndex = roomIndex
		newGamePlayer.UserIndex = userIndex
		newGamePlayer.ChairIndex = chairIndex
		newGamePlayer.BuyInLeft = buyInLeft
		newGamePlayer.LastActionDate = time.Now()

		if userNotFound == false {
			newGamePlayer.NickName = user.NickName
			newGamePlayer.Coin = user.Coin
		}

		dmap.AddGamePlayer(newGamePlayer)
	} else {
		return returnVal, errors.New("Exist GamePlayer")
	}
	return returnVal, nil
}

//ExitGame 게임을 나간다
func ExitGame(data map[string]interface{}) (interface{}, error) {
	var returnVal interface{}
	if data["roomIndex"] == nil || data["userIndex"] == nil {
		return returnVal, errors.New("Argument Error")
	}
	roomIndex, userIndex := int(data["roomIndex"].(float64)), int64(data["userIndex"].(float64))

	gamePlayer, err := dmap.GetGamePlayer(roomIndex, userIndex)

	if err != nil {
		return returnVal, err
	}
	dmap.RemoveGamePlayer(gamePlayer)
	return returnVal, nil
}
