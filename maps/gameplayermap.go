package maps

import (
	"errors"
	"holdempoker/models"
	"sync"
)

//GamePlayerMap 게임플레이어 데이터정보를 담는다.
type GamePlayerMap struct {
	data map[int][]models.GamePlayer
	lock *sync.Mutex
}

//GetGamePlayers 룸내 게임 플레이어들의 정보를 가져온다.
func (d *DataMap) GetGamePlayers(roomIndex int) ([]models.GamePlayer, error) {
	if gamePlayers, ok := d.gamePlayerMap.data[roomIndex]; ok {
		return gamePlayers, nil
	}
	return nil, errors.New("There is no GamePlayers")
}

//GetGamePlayer 게임플레이어 정보를 가져온다.
func (d *DataMap) GetGamePlayer(roomIndex int, userIndex int64) (models.GamePlayer, error) {
	if gamePlayers, err := d.GetGamePlayers(roomIndex); err == nil {
		if gamePlayers != nil {
			for i := 0; i < len(gamePlayers); i++ {
				if gamePlayers[i].UserIndex == userIndex {
					return gamePlayers[i], nil
				}
			}
		}
	}
	return models.GamePlayer{}, errors.New("There is no GamePlayes")
}

//ModifyGamePlayer 룸내 게임플레이어의 정보를 수정한다
func (d *DataMap) ModifyGamePlayer(gamePlayer models.GamePlayer) error {
	isUpdate := false
	d.gamePlayerMap.lock.Lock()
	if gamePlayers, err := d.GetGamePlayers(gamePlayer.RoomIndex); err == nil {
		if gamePlayers != nil {
			for i := 0; i < len(gamePlayers); i++ {
				if gamePlayers[i].UserIndex == gamePlayer.UserIndex {
					gamePlayers[i] = gamePlayer
					isUpdate = true
				}
			}
		} else {
			return err
		}

		if isUpdate == true {
			d.gamePlayerMap.data[gamePlayer.RoomIndex] = gamePlayers
		}
		d.gamePlayerMap.lock.Unlock()
	}
	return nil
}

//AddGamePlayer 게임플레이어를 추가한다.
func (d *DataMap) AddGamePlayer(gamePlayer models.GamePlayer) {
	d.gamePlayerMap.lock.Lock()
	gamePlayers, err := d.GetGamePlayers(gamePlayer.RoomIndex)
	if err != nil {
		gamePlayers = []models.GamePlayer{gamePlayer}
	} else {
		gamePlayers = append(gamePlayers, gamePlayer)
	}
	d.gamePlayerMap.data[gamePlayer.RoomIndex] = gamePlayers
	d.gamePlayerMap.lock.Unlock()
}

//AddGamePlayers 게임플레이어들을 단체로 넣는다.
func (d *DataMap) AddGamePlayers(roomIndex int, gamePlayers []models.GamePlayer) {
	d.gamePlayerMap.lock.Lock()
	if gamePlayers, ok := d.gamePlayerMap.data[roomIndex]; ok == false {
		d.gamePlayerMap.data[roomIndex] = gamePlayers
	}
	d.gamePlayerMap.lock.Unlock()
}

//RemoveGamePlayer 게임플레이어를 제거한다.
func (d *DataMap) RemoveGamePlayer(gamePlayer models.GamePlayer) {
	d.gamePlayerMap.lock.Lock()
	removeIndex := -1
	gamePlayers, err := d.GetGamePlayers(gamePlayer.RoomIndex)
	if err == nil {
		for i := 0; i < len(gamePlayers); i++ {
			if gamePlayers[i].UserIndex == gamePlayer.UserIndex {
				removeIndex = i
				break
			}
		}
		if removeIndex >= 0 {
			// 삭제로직
			copy(gamePlayers[removeIndex:], gamePlayers[removeIndex+1:])
			gamePlayers[len(gamePlayers)-1] = models.GamePlayer{}
			gamePlayers = gamePlayers[:len(gamePlayers)-1]
			d.gamePlayerMap.data[gamePlayer.RoomIndex] = gamePlayers
		}
	}
	d.gamePlayerMap.lock.Unlock()
}
