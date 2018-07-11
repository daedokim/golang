package routes

import (
	"errors"
	"holdempoker/models"
	"time"
)

//SetPlayerBet 게임을 나간다
func SetPlayerBet(data map[string]interface{}) (interface{}, error) {
	var returnVal interface{}
	if data["roomIndex"] == nil || data["userIndex"] == nil || data["betType"] == nil || data["callAmount"] == nil || data["betAmount"] == nil {
		return returnVal, errors.New("Argument Error")
	}
	roomIndex, userIndex := int(data["roomIndex"].(float64)), int64(data["userIndex"].(float64))
	betType, callAmount, betAmount := int(data["betType"].(float64)), int64(data["callAmount"].(float64)), int64(data["betAmount"].(float64))

	gamePlayer, err := dmap.GetGamePlayer(roomIndex, userIndex)

	if err != nil {
		return returnVal, err
	}

	room, err := dmap.GetRoom(roomIndex)
	if err != nil {
		return returnVal, err
	}

	isBlindBet := false
	totalAmount := callAmount + betAmount

	if betType == models.BetTypeAllin {
		totalAmount = betAmount
	}

	if betType == models.BetTypeBlind {
		betType = models.BetTypeRaise
		isBlindBet = true
	}

	if isBlindBet == true {
		gamePlayer.BetStatus = models.BetStatusBlindBetComplete
		if callAmount == 0 {
			room.StageBet = 0
			room.TotalBet = 0
			room.BetCount = 0
		}
	} else {
		gamePlayer.BetStatus = models.BetStatusBetComplete
		gamePlayer.LastActionDate = time.Now()
		gamePlayer.NoActionCount = 0
	}

	gamePlayer.BetCount++
	gamePlayer.LastBetType = betType
	gamePlayer.LastBet = totalAmount
	gamePlayer.LastCall = callAmount
	gamePlayer.LastRaise = betAmount
	gamePlayer.StageBet += totalAmount
	gamePlayer.TotalBet += totalAmount

	gamePlayer.BuyInLeft -= totalAmount
	gamePlayer.Stage = room.Stage

	if err := dmap.ModifyGamePlayer(gamePlayer); err != nil {
		return returnVal, err
	}

	if betType != models.BetTypeFold && betType != models.BetTypeCheck {
		room.BetCount++
	}
	room.StageBet += betAmount
	room.TotalBet += totalAmount

	room.LastRaise = betAmount
	room.LastBetType = betType

	if err := dmap.ModifyRoom(room); err != nil {
		return returnVal, err
	}

	if isBlindBet == false && (betType == models.BetTypeRaise || betType == models.BetTypeAllin && betAmount > 0) {
		if err := SetAnotherBetStatusReady(roomIndex, userIndex); err != nil {
			return returnVal, err
		}
	}

	return returnVal, nil
}

//SetAnotherBetStatusReady 다음 뱃 대상자가 있으면 해당 사용자세팅
func SetAnotherBetStatusReady(roomIndex int, userIndex int64) error {
	var gamePlayers []models.GamePlayer
	gamePlayers, err := dmap.GetGamePlayers(roomIndex)
	if err != nil {
		return err
	}

	for i := 0; i < len(gamePlayers); i++ {
		if gamePlayers[i].UserIndex == userIndex {
			continue
		}
		if gamePlayers[i].BetStatus == models.BetStatusBetComplete || gamePlayers[i].BetStatus == models.BetStatusBlindBetComplete {
			if gamePlayers[i].LastBetType != models.BetTypeFold && gamePlayers[i].LastBetType != models.BetTypeAllin {
				gamePlayers[i].BetStatus = models.BetStatusBetReady
				if err := dmap.ModifyGamePlayer(gamePlayers[i]); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
