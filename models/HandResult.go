package models

// HandResult is 족보 결과 정보.
type HandResult struct {
	cardType  int
	handType  int
	hands     [7]int
	kicks     [7]int
	madeCards [7]int
}

const (
	// CardTypeNone is NONE
	CardTypeNone = -1
	// CardTypeSpade is 스페이드
	CardTypeSpade = 1
	// CardTypeDiamond is 다이아몬드
	CardTypeDiamond = 2
	// CardTypeHeart is 하트
	CardTypeHeart = 3
	// CardTypeClover is 클로버
	CardTypeClover = 4

	// HandTypeNone is None
	HandTypeNone = -1
	// HandTypeRoyalStraightFlush is 로얄스트레이트 플러쉬
	HandTypeRoyalStraightFlush = 9
	// HandTypeStraightFlush is 스트레이트 플러쉬
	HandTypeStraightFlush = 8
	// HandTypePoker is 포카드
	HandTypePoker = 7
	// HandTypeFullHouse is 풀하우스
	HandTypeFullHouse = 6
	// HandTypeFlush is 플러쉬
	HandTypeFlush = 5
	// HandTypeStrait is 스트레이트
	HandTypeStrait = 4
	// HandTypeTriple is 트리플
	HandTypeTriple = 3
	// HandTypeTwoPair is 투페어
	HandTypeTwoPair = 2
	// HandTypeOnePair is 원페어
	HandTypeOnePair = 1
	// HandTypeTitle is 타이틀
	HandTypeTitle = 0
)

// NewHandResult is 족보 객체를 생성
func NewHandResult() *HandResult {
	r := new(HandResult)
	r.InitializeCards(&r.madeCards)
	r.InitializeCards(&r.hands)
	r.InitializeCards(&r.kicks)

	return r
}

// AddMadeCard is 메이드된 카드를 담는 메소드
func (r *HandResult) AddMadeCard(card int) {

	isFound := false

	for i := 0; i < len(r.madeCards); i++ {
		if card == r.madeCards[i] {
			isFound = true
			break
		}
	}

	// 메이드된 카드가 내 카드에 하나라도 발견되지 않는다면 메이드카드 배열에 넣는다.
	if isFound == false {
		for i := 0; i < len(r.madeCards); i++ {
			if r.madeCards[i] == -1 {
				r.madeCards[i] = card
				break
			}
		}
	}
}

// InitializeMadeCard is 메이드카드를 초기화 한다.
func (r *HandResult) InitializeMadeCard() {
	r.InitializeCards(&r.madeCards)
}

// InitializeCards is 카드를 초기화한다.
func (r HandResult) InitializeCards(cards *[7]int) {
	for i := 0; i < len(cards); i++ {
		cards[i] = -1
	}
}
