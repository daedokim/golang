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

//Sit 게임테이블에 참여한다.
func Sit(data map[string]interface{}) (interface{}, error) {
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

//StandUp 게임을 나간다
func StandUp(data map[string]interface{}) (interface{}, error) {
	var returnVal interface{}
	if data["roomIndex"] == nil || data["userIndex"] == nil {
		return returnVal, errors.New("Argument Error")
	}
	roomIndex, userIndex := int(data["roomIndex"].(float64)), int64(data["userIndex"].(float64))

	if gamePlayer, err := dmap.GetGamePlayer(roomIndex, userIndex); err == nil {
		if room, err := dmap.GetRoom(roomIndex); err == nil {
			if room.State != models.RoomStateWait {
				gamePlayer.State = models.GamePlayerStateStandWait
				dmap.ModifyGamePlayer(gamePlayer)
			} else {
				dmap.RemoveGamePlayer(gamePlayer)
			}

		} else {
			return returnVal, err
		}
	} else {
		return returnVal, err
	}

	return returnVal, nil
}
