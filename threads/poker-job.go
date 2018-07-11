package threads

import (
	"errors"
	"holdempoker/models"
	"holdempoker/utils"
	"time"

	"github.com/bradfitz/slice"
)

var helpers map[int]utils.CardSortingHelper

//PokerJob 포커 스케줄러 작동
func PokerJob() {
	ticker := time.NewTicker(500 * time.Millisecond)
	helpers = make(map[int]utils.CardSortingHelper)

	var timeTick int64
	go func() {
		for range ticker.C {
			diffTime := utils.GetCurrentTick() - timeTick
			timeTick = utils.GetCurrentTick()
			rooms := dmap.GetRooms()

			for i := 0; i < len(rooms); i++ {
				if _, ok := helpers[rooms[i].RoomIndex]; ok == false {
					helpers[rooms[i].RoomIndex] = utils.CardSortingHelper{}
				}
				CheckWaitTime(&rooms[i], int(diffTime))

				switch rooms[i].State {
				case models.RoomStateWait:
					GotoReady(rooms[i])
				case models.RoomStateReady:
					GotoSetting(rooms[i])
				case models.RoomStateSetting:
					GotoPlay(rooms[i])
				case models.RoomStatePlaying:
					SetPlaying(rooms[i])
				}
			}

			ClearNotUseHelper(rooms)
		}
	}()
}

//ClearNotUseHelper 불필요한 카드 핼퍼 처리
func ClearNotUseHelper(rooms []models.Room) {
	for key := range helpers {
		if _, ok := helpers[key]; ok == false {
			delete(helpers, key)
		}
	}
}

//GotoReady 대기중인 방일때
func GotoReady(room models.Room) {
	playerList, err := dmap.GetGamePlayers(room.RoomIndex)

	if err == nil {
		if len(playerList) >= 2 {
			room.State = models.RoomStateReady
			room.Stage = 1
			room.WaitTimeout = models.WaitTimeoutForReady

			if helper, ok := helpers[room.RoomIndex]; ok == true {
				helper.Initialize()
				room.Card1 = helper.Pop()
				room.Card2 = helper.Pop()
				room.Card3 = helper.Pop()
				room.Card4 = helper.Pop()
				room.Card5 = helper.Pop()
			}
			dmap.ModifyRoom(room)
		}
	}
}

//GotoSetting 대기중인 방일때
func GotoSetting(room models.Room) {

	playerList, _ := GetJoinedPlayers(room.RoomIndex)

	if room.WaitTimeout <= 0 {
		if len(playerList) >= 2 {
			room.State = models.RoomStateSetting
			room.Stage = 2
			room.WaitTimeout = models.WaitTimeoutForSetting
			room.DealerChairIndex = GetNearChairIndex(room.DealerChairIndex, playerList)
			room.OwnerIndex = GetNearChairIndex(room.DealerChairIndex, playerList)

			//room.lastbet = 0;
			room.StageBet = 0
			room.LastRaise = 0
			room.CurrentUserIndex = GetUserIndexByOwnerIndex(playerList, room.OwnerIndex)
			room.CurrentOrderNo = GetOrderByOwnerIndex(playerList, room.OwnerIndex)

			playerList = SetPlayersCard(helpers[room.RoomIndex], playerList)
			playerList = SetPlayerOrderNo(room.OwnerIndex, playerList)

			for i := 0; i < len(playerList); i++ {
				InitGamePlayerMember(&playerList[i])
				dmap.ModifyGamePlayer(playerList[i])
			}

		} else {
			room.State = models.RoomStateWait
			room.Stage = 0
		}
		dmap.ModifyRoom(room)
	}
}

//GotoPlay 대기중인 방일때
func GotoPlay(room models.Room) {
	if room.WaitTimeout <= 0 {
		room.State = models.RoomStatePlaying
		room.Stage = 3
		room.WaitTimeout = models.WaitTimeoutForGamePlayer
		dmap.ModifyRoom(room)
	}
}

//SetPlaying 대기중인 방일때
func SetPlaying(room models.Room) {
	stageSet := room.Stage % 3

	if room.Stage >= 3 && room.Stage < 14 {
		switch stageSet {
		case 0:
			CheckBetStatus(room)
		case 1:
			CheckWinner(room)
		case 2:
			CheckSetting(room)
		}
		CheckGameStatus(room)
	}
	if room.Stage == 14 || room.Stage == 15 {
		if room.WaitTimeout <= 0 {
			GotoFinish(room)
		}
	} else if room.Stage == 17 {
		if room.WaitTimeout <= 0 {
			GotoInitialize(room)
		}
	}
}

//CheckWaitTime 대기시간 체크
func CheckWaitTime(room *models.Room, diffTime int) {

	if room.WaitTimeout > 0 {
		room.WaitTimeout -= diffTime

		if room.WaitTimeout < 0 {
			room.WaitTimeout = 0
		}
	}
}

//GetJoinedPlayers 게임을 진행하는 유저 목록 가져오기
func GetJoinedPlayers(roomIndex int) ([]models.GamePlayer, error) {
	if gamePlayers, err := dmap.GetGamePlayers(roomIndex); err == nil {
		joinedPlayers := make([]models.GamePlayer, 0, 0)

		for i := 0; i < len(gamePlayers); i++ {
			if gamePlayers[i].State == models.GamePlayerStateSitWait || gamePlayers[i].State == models.GamePlayerStatePlay || gamePlayers[i].State == models.GamePlayerStateStandWait {
				joinedPlayers = append(joinedPlayers, gamePlayers[i])
			}
		}
		slice.Sort(gamePlayers[:], func(i, j int) bool {
			return gamePlayers[i].ChairIndex < gamePlayers[j].ChairIndex
		})

		return gamePlayers, nil
	}
	return nil, errors.New("There is no JoinedGamePlayers")
}

//GetNearChairIndex 지정한 체어 인덱스의 인접한 다음 체어인덱스를 가져온다.
func GetNearChairIndex(dealerChairIndex int, playerList []models.GamePlayer) int {
	chairIndex := -1
	tempChairIndex, selectedChairIndex, minDiff := 0, 0, 999

	for i := 0; i < len(playerList); i++ {
		chairIndex = playerList[i].ChairIndex

		// 이전 딜러 체어 인덱스에 가장 가까운 체어 인덱스를 찾는다.
		if chairIndex < dealerChairIndex {
			tempChairIndex = chairIndex + models.MaxGamePlayerCount
		} else {
			tempChairIndex = chairIndex
		}
		if chairIndex != dealerChairIndex && tempChairIndex-dealerChairIndex < minDiff {
			minDiff = tempChairIndex - dealerChairIndex
			selectedChairIndex = chairIndex
		}
	}
	return selectedChairIndex
}

//GetUserIndexByOwnerIndex 오너인덱스로 유저인덱스를 가져온다.
func GetUserIndexByOwnerIndex(playerList []models.GamePlayer, ownerIndex int) int64 {
	var userIndex int64
	for i := 0; i < len(playerList); i++ {
		if playerList[i].ChairIndex == ownerIndex {
			userIndex = playerList[i].UserIndex
			break
		}
	}
	return userIndex
}

// GetOrderByOwnerIndex 오너인덱스로 정령번호를 가져온다.
func GetOrderByOwnerIndex(playerList []models.GamePlayer, ownerIndex int) int {
	orderNo := 0
	for i := 0; i < len(playerList); i++ {
		if playerList[i].ChairIndex == ownerIndex {
			orderNo = playerList[i].OrderNo
			break
		}
	}
	return orderNo
}

// GetUserIndexByOrderNo 정렬번호로 유저인덱스를 가져온다.
func GetUserIndexByOrderNo(roomIndex int, orderNo int) int64 {
	var userIndex int64
	if playerList, err := GetJoinedPlayers(roomIndex); err == nil {
		for i := 0; i < len(playerList); i++ {
			if playerList[i].OrderNo == orderNo {
				userIndex = playerList[i].UserIndex
				break
			}
		}
	}
	return userIndex
}

//SetPlayersCard 플레이어들에게 카드를 세팅한다.
func SetPlayersCard(helper utils.CardSortingHelper, playerList []models.GamePlayer) []models.GamePlayer {
	for i := 0; i < len(playerList); i++ {
		playerList[i].Card1 = helper.Pop()
		playerList[i].Card2 = helper.Pop()
	}

	return playerList
}

//SetPlayerOrderNo 플레이의 정렬순서를 세팅한다.
func SetPlayerOrderNo(ownerIndex int, playerList []models.GamePlayer) []models.GamePlayer {
	ownerNo := 0
	for i := 0; i < len(playerList); i++ {
		if playerList[i].ChairIndex == ownerIndex {
			ownerNo = i
			break
		}
	}
	for i := 0; i < len(playerList); i++ {
		if i < ownerNo {
			playerList[i].OrderNo = len(playerList) - ownerNo + i
		} else {
			playerList[i].OrderNo = i - ownerNo
		}
	}

	return playerList
}

//InitGamePlayerMember 멤버들의 게임데이터 초기화
func InitGamePlayerMember(gamePlayer *models.GamePlayer) {
	gamePlayer.State = models.GamePlayerStatePlay
	gamePlayer.BetStatus = models.BetStatusBetReady
	gamePlayer.LastBetType = 0
	gamePlayer.StageBet = 0
}

//CheckBetStatus 뱃상태를 체크한다.
func CheckBetStatus(room models.Room) {
	var currentUserIndex int64
	currentOrderNo := -1

	if gamePlayer, err := dmap.GetGamePlayer(room.RoomIndex, room.CurrentUserIndex); err == nil {
		isCompleteBetUser := false

		if (gamePlayer.BetStatus == models.BetStatusBlindBetComplete || gamePlayer.BetStatus == models.BetStatusBetComplete) && (gamePlayer.State == models.GamePlayerStatePlay || gamePlayer.State == models.GamePlayerStateStandWait) {
			isCompleteBetUser = true
		}

		if room.WaitTimeout <= 0 || isCompleteBetUser == true {
			if gamePlayer.BetStatus == models.BetStatusBetReady {
				gamePlayer.BetType = models.BetTypeFold
				gamePlayer.BetStatus = models.BetStatusBetComplete

				gamePlayer.LastBetType = gamePlayer.BetType
				gamePlayer.Stage = room.Stage
				dmap.ModifyGamePlayer(gamePlayer)
			}

			// 대기하고있는 사용자가 하나도 없다면 다음 스테이지로 이동 한다.
			if GetReadyCount(room.RoomIndex) == 0 {
				room.Stage++
				room.BetCount = 0
				room.WaitTimeout = models.WaitTimeoutForSetting
				currentOrderNo = 0
				currentUserIndex = 0
				room.LastRaise = 0
				dmap.ModifyRoom(room)
			} else {
				room.WaitTimeout = models.WaitTimeoutForGamePlayer
				currentOrderNo = GetNextOrderNo(room.RoomIndex, room.CurrentOrderNo)
				currentUserIndex = GetUserIndexByOrderNo(room.RoomIndex, currentOrderNo)

				if currentUserIndex > 0 {
					room.CurrentUserIndex = currentUserIndex
					room.CurrentOrderNo = currentOrderNo
					dmap.ModifyRoom(room)
				}
			}
		}
	}
}

//GetReadyCount 대기중인 사용자들 카운트를 가져온다.
func GetReadyCount(roomIndex int) int {
	count := 0
	if playerList, err := GetJoinedPlayers(roomIndex); err == nil {
		for i := 0; i < len(playerList); i++ {
			if playerList[i].BetStatus == models.BetStatusBetReady && playerList[i].LastBetType != models.BetTypeFold && playerList[i].LastBetType != models.BetTypeAllin && (playerList[i].State == models.GamePlayerStatePlay || playerList[i].State == models.GamePlayerStateStandWait) {
				count++
			}
		}
	}
	return count
}

//GetNextOrderNo 다음 정렬순서를 가져온다.
func GetNextOrderNo(roomIndex int, orderNo int) int {
	currentOrderNo := -1
	if gamePlayers, err := GetJoinedPlayers(roomIndex); err == nil {

		slice.Sort(gamePlayers[:], func(i, j int) bool {
			return gamePlayers[i].OrderNo < gamePlayers[j].OrderNo
		})

		orderNo = (orderNo + 1) % len(gamePlayers)

		for i := 0; i < len(gamePlayers); i++ {
			if orderNo == gamePlayers[i].OrderNo {
				if gamePlayers[i].BetStatus == models.BetStatusBetReady && gamePlayers[i].LastBetType != models.BetTypeFold && gamePlayers[i].LastBetType != models.BetTypeAllin && (gamePlayers[i].State == models.GamePlayerStatePlay || gamePlayers[i].State == models.GamePlayerStateStandWait) {
					currentOrderNo = orderNo
				} else {
					currentOrderNo = GetNextOrderNo(roomIndex, orderNo)
				}
				break
			}
		}
	}
	return currentOrderNo
}

//CheckSetting 세팅상태를 체크한다.
func CheckSetting(room models.Room) {
	if room.WaitTimeout <= 0 {
		room.Stage++

		room.LastRaise = room.MinbetAmount
		room.WaitTimeout = models.WaitTimeoutForGamePlayer

		room.CurrentOrderNo = GetFirstOrderNo(room.RoomIndex)
		room.CurrentUserIndex = GetUserIndexByOrderNo(room.RoomIndex, room.CurrentOrderNo)

		dmap.ModifyRoom(room)
		ClearGamePlayerByStage(room.RoomIndex, room.Stage)
	}
}

//GetFirstOrderNo 첫 정렬순서를 가져온다.
func GetFirstOrderNo(roomIndex int) int {
	orderNo := 0
	if gamePlayers, err := GetJoinedPlayers(roomIndex); err == nil {
		slice.Sort(gamePlayers[:], func(i, j int) bool {
			return gamePlayers[i].OrderNo < gamePlayers[j].OrderNo
		})

		for i := 0; i < len(gamePlayers); i++ {
			if gamePlayers[i].LastBetType != models.BetTypeAllin && gamePlayers[i].LastBetType != models.BetTypeFold {
				orderNo = gamePlayers[i].OrderNo
				break
			}
		}
	}
	return orderNo
}

//ClearGamePlayerByStage 플레이어들의 각각 스테이지를 정리한다.
func ClearGamePlayerByStage(roomIndex int, stage int) {
	if playerList, err := GetJoinedPlayers(roomIndex); err == nil {
		for i := 0; i < len(playerList); i++ {
			if (playerList[i].BetStatus == models.BetStatusBetComplete || playerList[i].BetStatus == models.BetStatusBlindBetComplete) && playerList[i].LastBetType != models.BetTypeFold && playerList[i].LastBetType != models.BetTypeAllin && (playerList[i].State == models.GamePlayerStatePlay || playerList[i].State == models.GamePlayerStateStandWait) {
				playerList[i].BetStatus = models.BetStatusBetReady
				playerList[i].LastBetType = 0
				playerList[i].StageBet = 0

				dmap.ModifyGamePlayer(playerList[i])
			}
		}
	}
}

//CheckGameStatus 게임상태를 체크한다.
func CheckGameStatus(room models.Room) {
	playerCount := 0
	var userIndex int64

	if playerList, err := GetJoinedPlayers(room.RoomIndex); err == nil {
		for i := 0; i < len(playerList); i++ {
			if playerList[i].LastBetType != models.BetTypeFold && (playerList[i].State == models.GamePlayerStatePlay || playerList[i].State == models.GamePlayerStateStandWait) {
				userIndex = playerList[i].UserIndex
				playerCount++
			}
		}

		if playerCount == 0 {
			room.WaitTimeout = models.WaitTimeoutForReady
			room.Stage = -1
			room.OwnerIndex = -1
			room.WinnerUserIndex = -1
			room.CurrentUserIndex = 0
			room.CurrentOrderNo = -1
			dmap.ModifyRoom(room)
		} else if playerCount == 1 {
			room.WaitTimeout = models.WaitTimeoutForReady
			room.Stage = 15
			room.WinnerUserIndex = userIndex
			room.CurrentUserIndex = 0
			room.CurrentOrderNo = -1
			dmap.ModifyRoom(room)
		}

	}
}

// GotoInitialize 초기화로 이동
func GotoInitialize(room models.Room) {
	room.State = models.RoomStateWait
	room.Stage = 0
	room.TotalBet = 0
	dmap.ModifyRoom(room)
}

// GotoFinish 최종평가로 이동
func GotoFinish(room models.Room) {
	room.WaitTimeout = models.WaitTimeoutForInit
	room.Stage = 17

	if playerList, err := GetJoinedPlayers(room.RoomIndex); err == nil {
		for i := 0; i < len(playerList); i++ {
			if playerList[i].State == models.GamePlayerStateStandWait {
				dmap.RemoveGamePlayer(playerList[i])
			}
		}
	}

	dmap.ModifyRoom(room)
}

//CheckWinner 승자를 체크한다.
func CheckWinner(room models.Room) {
	if room.Stage == 13 {
		if playerList, err := GetJoinedPlayers(room.RoomIndex); err == nil {
			validator := utils.WinnerValidator{}

			for i := 0; i < len(playerList); i++ {
				if playerList[i].State == models.GamePlayerStateStandWait {
					result := validator.GetResult([]int{room.Card1, room.Card2, room.Card3, room.Card4, room.Card5, playerList[i].Card1, playerList[i].Card2})
					playerList[i].Result = result
					dmap.ModifyGamePlayer(playerList[i])
				}
			}
			room.WinnerUserIndex = validator.GetWinner(playerList)
			room.WaitTimeout = models.WaitTimeoutForReady
			room.CurrentUserIndex = 0
			room.CurrentOrderNo = -1
			room.Stage = 14
			dmap.ModifyRoom(room)
		}
	} else {
		room.Stage++
		room.WaitTimeout = models.WaitTimeoutForSetting
		room.StageBet = 0
		room.CurrentOrderNo = 0
		room.CurrentUserIndex = 0

		dmap.ModifyRoom(room)
	}
}
