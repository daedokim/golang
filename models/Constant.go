package models

const (
	// RoomStateWait is 대기
	RoomStateWait = -1
	// RoomStateReady is 게임준비
	RoomStateReady = 0
	// RoomStateSetting is 테이블 세팅
	RoomStateSetting = 1
	// RoomStatePlaying is 게임중
	RoomStatePlaying = 2

	// StagePreFlop is 프리플롭
	StagePreFlop = 3
	// StageFlop is 플롭
	StageFlop = 6
	// StageTurn is 턴
	StageTurn = 9
	// StageRiver is 리버
	StageRiver = 12

	// GamePlayerStateStand is 게임을 안하는 상태
	GamePlayerStateStand = 0
	// GamePlayerStatePlay is 게임중인 상태
	GamePlayerStatePlay = 1
	// GamePlayerStateSitWait is 자리에 앉아 대기중인 상태
	GamePlayerStateSitWait = 2
	// GamePlayerStateStandWait is 일어나기 예약 상태
	GamePlayerStateStandWait = 3

	// BetStatusBetReady is 벳 준비 상태
	BetStatusBetReady = 0
	// BetStatusBetComplete is 벳 완료 상태
	BetStatusBetComplete = 1
	// BetStatusAllBetComplete is 모든 플레이어 벳완료 상태
	BetStatusAllBetComplete = 2
	// BetStatusBlindBetComplete is 블라인드 벳 완료 상태
	BetStatusBlindBetComplete = 3

	// BetTypeCheck is 뱃타입:체크
	BetTypeCheck = 1
	// BetTypeCall is 뱃타입:콜
	BetTypeCall = 2
	// BetTypeBlind is 뱃타입:블라인드
	BetTypeBlind = 3
	// BetTypeRaise is 뱃타입:레이스
	BetTypeRaise = 4
	// BetTypeAllin is 뱃타입:올인
	BetTypeAllin = 5
	// BetTypeFold is 뱃타입:폴드
	BetTypeFold = 6
)
