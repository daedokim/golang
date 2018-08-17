package threads

import (
	database "holdempoker/db"
	"holdempoker/maps"
	"holdempoker/models"

	"github.com/jinzhu/gorm"
)

//Thread 쓰레드
type Thread struct {
}

var db *gorm.DB
var dmap *maps.DataMap

//Job 스케줄러 시작
func (t *Thread) Job() {

	PokerJob()
	//DataJob()
}

//InitJob 서버가 초기화될때 DB데이터와 동기화를 한다.
func (t *Thread) InitJob() {

	db = database.GetInstance().GetDB()
	dmap = maps.GetInstance()

	var rooms []models.Room
	var gamePlayers []models.GamePlayer

	if notFound := db.Find(&rooms).RecordNotFound(); notFound == false {
		for i := 0; i < len(rooms); i++ {
			if _, err := dmap.GetRoom(rooms[i].RoomIndex); err != nil {
				dmap.AddRoom(rooms[i])
			}
		}
	}

	if notFound := db.Find(&gamePlayers).RecordNotFound(); notFound == false {
		for i := 0; i < len(gamePlayers); i++ {
			if _, err := dmap.GetGamePlayer(gamePlayers[i].RoomIndex, gamePlayers[i].UserIndex); err != nil {
				dmap.AddGamePlayer(gamePlayers[i])
			}
		}
	}

}
