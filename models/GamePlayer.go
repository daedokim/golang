package models

import (
	"time"
)

// GamePlayer is 게임플레이어 정보
type GamePlayer struct {
	state      int
	chairIndex int
	roomIndex  int
	round      int

	card1 int
	card2 int

	buyInLeft      int64
	orderNo        int
	stage          int
	betStatus      int
	betType        int
	lastBetType    int
	betCount       int
	lastBet        int64
	lastCall       int64
	lastRaise      int64
	totalBet       int64
	stageBet       int64
	lastActionDate time.Time
	noActionCount  int
	coin           int64
	nickName       string
	userIndex      int64
	result         HandResult
}
