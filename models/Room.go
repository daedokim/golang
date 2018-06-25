package models

// Room is 게임룸 정보
type Room struct {
	Round            int   `json:"round"`
	Card1            int   `json:"card1"`
	Card2            int   `json:"card2"`
	Card3            int   `json:"card3"`
	Card4            int   `json:"card4"`
	Card5            int   `json:"card5"`
	Lastbet          int64 `json:"lastbet"`
	LastRaise        int64 `json:"lastRaise"`
	WinnerUserIndex  int64 `json:"winnerUserIndex"`
	CurrentUserIndex int64 `json:"currentUserIndex"`
	DealerChairIndex int   `json:"dealerChairIndex"`
	Owner            int64 `json:"owner"`
	TotalBet         int64 `json:"totalBet"`
	LastbetType      int   `json:"lastbetType"`
	LastBet          int64 `json:"lastBet"`
	StageBet         int64 `json:"stageBet"`
	BuyInMin         int64 `json:"buyInMin"`
	BuyInMax         int64 `json:"buyInMax"`
	Stage            int   `json:"stage"`
	Betfinished      int   `json:"betfinished"`
	BetCount         int   `json:"betCount"`
	CurrentOrderNo   int   `json:"currentOrderNo"`
	WaitTimeout      int   `json:"waitTimeout"`
	MinbetAmount     int64 `json:"minbetAmount"`
	RoomState        int   `json:"roomState"`
}

// TableName is 테이블 이름
func (Room) TableName() string {
	return "ROOM"
}
