package maps

import (
	"errors"
	"holdempoker/models"
	"sync"
)

//RoomMap 정보
type RoomMap struct {
	data map[int]models.Room
	lock *sync.Mutex
}

//GetRooms 룸 목록을 가져온다.
func (d *DataMap) GetRooms() []models.Room {
	keys := make([]models.Room, 0, 3)
	for k := range d.roomMap.data {
		keys = append(keys, d.roomMap.data[k])
	}
	return keys
}

//GetRoom 룸정보를 가져온다.
func (d *DataMap) GetRoom(roomIndex int) (models.Room, error) {
	if room, ok := d.roomMap.data[roomIndex]; ok {
		return room, nil
	}
	return models.Room{}, errors.New("Not found room")
}

//ModifyRoom 룸정보를 수정한다..
func (d *DataMap) ModifyRoom(room models.Room) error {
	d.roomMap.lock.Lock()
	r, err := d.GetRoom(room.RoomIndex)
	if err != nil {
		d.roomMap.lock.Unlock()
		return err
	}
	r.Update(room)
	d.roomMap.lock.Unlock()
	return nil
}

//AddRoom 룸을 추가 한다.
func (d *DataMap) AddRoom(room models.Room) error {
	d.roomMap.lock.Lock()
	_, err := d.GetRoom(room.RoomIndex)
	if err == nil {
		d.roomMap.lock.Unlock()
		return errors.New("Aready has room")
	}
	d.roomMap.data[room.RoomIndex] = room
	d.roomMap.lock.Unlock()
	return nil
}

//RemoveRoom room을 삭제한다.
func (d *DataMap) RemoveRoom(room models.Room) error {
	d.roomMap.lock.Lock()
	delete(d.roomMap.data, 777)
	d.roomMap.lock.Unlock()
	return nil
}
