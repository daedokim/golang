package models

import (
	"time"
)

// GamePlayer is 게임플레이어 정보
type GamePlayer struct {
	State          int        `json:"state"`
	ChairIndex     int        `json:"chairIndex"`
	RoomIndex      int        `json:"roomIndex"`
	Round          int        `json:"round"`
	Card1          int        `json:"card1"`
	Card2          int        `json:"card2"`
	BuyInLeft      int64      `json:"buyInLeft"`
	OrderNo        int        `json:"orderNo"`
	Stage          int        `json:"stage"`
	BetStatus      int        `json:"betStatus"`
	BetType        int        `json:"betType"`
	LastBetType    int        `json:"lastBetType"`
	BetCount       int        `json:"betCount"`
	LastBet        int64      `json:"lastBet"`
	LastCall       int64      `json:"lastCall"`
	LastRaise      int64      `json:"lastRaise"`
	TotalBet       int64      `json:"totalBet"`
	StageBet       int64      `json:"stageBet"`
	LastActionDate time.Time  `json:"lastActionDate"`
	NoActionCount  int        `json:"noActionCount"`
	Coin           int64      `json:"coin"`
	NickName       string     `json:"nickName"`
	UserIndex      int64      `json:"userIndex"`
	Result         HandResult `json:"result"`
}

// TableName 테이블 이름
func (GamePlayer) TableName() string {
	return "tbl_game_player"
}
