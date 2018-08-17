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
					for j := 0; j < len(gamePlayers); j++ {
						if notFound := db.Where("room_index = ? AND user_index = ?", gamePlayers[j].RoomIndex, gamePlayers[j].UserIndex).First(&gamePlayer).RecordNotFound(); notFound == true {
							session := db.Begin()
							if err := session.Create(gamePlayers[j]).Error; err != nil {
								session.Rollback()
							}
							session.Commit()
						} else {
							session := db.Begin()
							if err := session.Update(gamePlayers[j]).Error; err != nil {
								session.Rollback()
							}
							session.Commit()
						}

					}
				}
			}

			var noActionGamePlayer []models.GamePlayer
			if notFound := db.Where("TIME_TO_SEC(timediff(now(),  last_action_date)) > ?", 60*5).Find(&noActionGamePlayer).RecordNotFound(); notFound == false {
				for j := 0; j < len(noActionGamePlayer); j++ {
					if _, err := dmap.GetGamePlayer(noActionGamePlayer[j].RoomIndex, noActionGamePlayer[j].UserIndex); err == nil {
						session := db.Begin()
						if err := db.Delete(noActionGamePlayer[j]); err != nil {
							session.Rollback()
						}
						session.Commit()
						dmap.RemoveGamePlayer(noActionGamePlayer[j])
					}
				}
			}

		}
	}()
}
