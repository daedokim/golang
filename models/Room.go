package models

// Room is 게임룸 정보
type Room struct {
	RoomIndex        int   `json:"roomIndex" gorm:"primary_key"`
	State            int   `json:"state"`
	Round            int   `json:"round"`
	Card1            int   `json:"card1"`
	Card2            int   `json:"card2"`
	Card3            int   `json:"card3"`
	Card4            int   `json:"card4"`
	Card5            int   `json:"card5"`
	LastBet          int64 `json:"lastBet"`
	LastRaise        int64 `json:"lastRaise"`
	WinnerUserIndex  int64 `json:"winnerUserIndex"`
	CurrentUserIndex int64 `json:"currentUserIndex"`
	DealerChairIndex int   `json:"dealerChairIndex"`
	OwnerIndex       int64 `json:"ownerIndex"`
	TotalBet         int64 `json:"totalBet"`
	StageBet         int64 `json:"stageBet"`
	LastBetType      int   `json:"lastBetType"`
	BuyInMin         int64 `json:"buyInMin"`
	BuyInMax         int64 `json:"buyInMax"`
	Stage            int   `json:"stage"`
	BetFinished      int   `json:"betFinished"`
	BetCount         int   `json:"betCount"`
	CurrentOrderNo   int   `json:"currentOrderNo"`
	MinbetAmount     int64 `json:"minbetAmount"`
	WaitTimeout      int   `json:"waitTimeout"`
}

// TableName 테이블 이름
func (Room) TableName() string {
	return "tbl_room"
}
