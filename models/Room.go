package models

// Room isx임룸 정보
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

// Update 값을 업데이트
func (r *Room) Update(room Room) {
	r.RoomIndex = room.RoomIndex
	r.State = room.State
	r.Round = room.Round
	r.Card1 = room.Card1
	r.Card2 = room.Card2
	r.Card3 = room.Card3
	r.Card4 = room.Card4
	r.Card5 = room.Card5
	r.LastBet = room.LastBet
	r.LastRaise = room.LastRaise
	r.WinnerUserIndex = room.WinnerUserIndex
	r.CurrentUserIndex = room.CurrentUserIndex
	r.DealerChairIndex = room.DealerChairIndex
	r.OwnerIndex = room.OwnerIndex
	r.TotalBet = room.TotalBet
	r.StageBet = room.StageBet
	r.LastBetType = room.LastBetType
	r.BuyInMin = room.BuyInMin
	r.BuyInMax = room.BuyInMax
	r.Stage = room.Stage
	r.BetFinished = room.BetFinished
	r.BetCount = room.BetCount
	r.CurrentOrderNo = room.CurrentOrderNo
	r.MinbetAmount = room.MinbetAmount
	r.WaitTimeout = room.WaitTimeout
}
