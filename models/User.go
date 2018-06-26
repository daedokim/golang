package models

import (
	"time"
)

// User 유저정보
type User struct {
	UserIndex int64 `gorm:"primary_key"`
	Coin      int64
	Nickname  string
	LoginDate time.Time
	WriteDate time.Time
	UserID    string
	Passwd    string
}

// TableName 유저테이블명
func (u User) TableName() string {
	return "tbl_user"
}

// UserAuth 유저 권한 정보
type UserAuth struct {
	UserIndex int64
	UID       string
	OsType    int
}

// TableName 유저테이블명
func (u UserAuth) TableName() string {
	return "tbl_user_auth"
}
