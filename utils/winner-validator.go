package utils

import "holdempoker/models"

//WinnerValidator 승자계산
type WinnerValidator struct {
	handUtil PokerHandUtil
}

//GetResult 결과를 조회한다.
func (w *WinnerValidator) GetResult(cards []int) models.HandResult {
	w.handUtil = PokerHandUtil{}
	return w.handUtil.CheckHands(cards)
}

//GetWinner 승자를 조회한다.
func (w *WinnerValidator) GetWinner(playerList []models.GamePlayer) int64 {
	var winnerList []models.GamePlayer
	maxHand := -1
	currentHand := -1
	var userIndex int64

	for i := 0; i < len(playerList); i++ {
		currentHand = playerList[i].Result.HandType
		if maxHand < currentHand {
			maxHand = currentHand
			winnerList = make([]models.GamePlayer, 0, 0)
			winnerList = append(winnerList, playerList[i])
		} else if maxHand == currentHand {
			winnerList = append(winnerList, playerList[i])
		}
	}
	if winnerList != nil {
		if len(winnerList) == 1 {
			userIndex = winnerList[0].UserIndex
		} else if len(winnerList) > 1 {
			userIndex = GetTiedWinner(winnerList)
			if userIndex == 0 {
				userIndex = GetKickWinner(winnerList)
			}
		}
	}
	return userIndex
}

//GetTiedWinner 같은 족보의 승리자들중에서 승자를 조회한다.
func GetTiedWinner(resultList []models.GamePlayer) int64 {
	var userIndex int64
	maxValue := 0
	for i := 0; i < len(resultList); i++ {
		for j := 0; j < len(resultList[i].Result.Hands); j++ {
			if resultList[i].Result.Hands[j] > maxValue {
				maxValue = resultList[i].Result.Hands[j]
				userIndex = resultList[i].UserIndex
			}
		}
	}

	return userIndex
}

//GetKickWinner 같은 족보의 승리자들중에서 승자를 조회한다.
func GetKickWinner(resultList []models.GamePlayer) int64 {
	var userIndex int64
	maxValue := 0
	for i := 0; i < len(resultList); i++ {
		for j := 0; j < len(resultList[i].Result.Kicks); j++ {
			if resultList[i].Result.Kicks[j] > maxValue {
				maxValue = resultList[i].Result.Kicks[j]
				userIndex = resultList[i].UserIndex
			}
		}
	}

	return userIndex
}
