package maps

import (
	"holdempoker/models"
	"sync"
)

//DataMap 데이터 맵
type DataMap struct {
	roomMap       *RoomMap
	gamePlayerMap *GamePlayerMap
}

var once sync.Once
var instance *DataMap

// GetInstance is 인스턴스 생성
func GetInstance() *DataMap {
	once.Do(func() {
		instance = &DataMap{}
		instance.roomMap = &RoomMap{data: make(map[int]models.Room), lock: &sync.Mutex{}}
		instance.gamePlayerMap = &GamePlayerMap{data: make(map[int][]models.GamePlayer), lock: &sync.Mutex{}}

	})
	return instance
}
