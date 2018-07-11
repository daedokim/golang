package threads

import (
	"holdempoker/models"
	"time"
)

//DataJob DB데이터와 서버 메모리 싱크 담당
func DataJob() {
	ticker := time.NewTicker(5000 * time.Millisecond)

	var room models.Room
	var gamePlayer models.GamePlayer
	go func() {
		for range ticker.C {
			rooms := dmap.GetRooms()
			for i := 0; i < len(rooms); i++ {
				if notFound := db.Where("room_index = ?", rooms[i].RoomIndex).First(&room).RecordNotFound(); notFound == true {
					session := db.Begin()
					if err := session.Create(rooms[i]).Error; err != nil {
						session.Rollback()
					}
					session.Commit()
				} else {
					session := db.Begin()
					if err := session.Update(rooms[i]).Error; err != nil {
						session.Rollback()
					}
					session.Commit()
				}

				if gamePlayers, err := dmap.GetGamePlayers(rooms[i].RoomIndex); err == nil {
					for i := 0; i < len(gamePlayers); i++ {
						if notFound := db.Where("room_index = ?", gamePlayers[i].RoomIndex).First(&gamePlayer).RecordNotFound(); notFound == true {
							session := db.Begin()
							if err := session.Create(gamePlayers[i]).Error; err != nil {
								session.Rollback()
							}
							session.Commit()
						} else {
							session := db.Begin()
							if err := session.Update(gamePlayers[i]).Error; err != nil {
								session.Rollback()
							}
							session.Commit()
						}
					}
				}

			}
		}
	}()
}
