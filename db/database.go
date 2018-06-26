package db

import (
	"fmt"
	"sync"

	. "holdempoker/config"

	"github.com/jinzhu/gorm"
)

type Database struct {
}

var instance *Database
var once sync.Once

// GetDBInstance is 인스턴스 생성
func GetDBInstance() *Database {
	once.Do(func() {
		instance = &Database{}
	})
	return instance
}

//DB Start
var DB *gorm.DB

//OpenDatabase start
func (d *Database) OpenDatabase(Conf *Config) {
	var err error
	DB, err = gorm.Open(Conf.Database.Driver, Conf.Database.Connection)

	if err != nil {
		panic(err)
	}
	fmt.Println("--- Open database ---")
	DB.LogMode(Conf.Debug)
	DB.SingularTable(true)
}

// GetDB 는 gormDB를 가져오기위한 메소드
func (d *Database) GetDB() *gorm.DB {
	return DB
}
