package utils

import (
	"holdempoker/models"
	"strconv"
	"strings"

	"github.com/bradfitz/slice"
)

//PokerHandUtil 포커 족보 유틸
type PokerHandUtil struct {
	hands []interface{}
}

// CheckHands 족보를 체크한다.
func (p *PokerHandUtil) CheckHands(cards []int) models.HandResult {

	var result models.HandResult
	var funcRef interface{}
	p.SetHands()

	for i := 0; i < len(p.hands); i++ {
		funcRef = p.hands[i]
		funcRef.(func([]int) models.HandResult)(cards)
	}
	return result
}

//SetHands 족보체크 함수 추가
func (p *PokerHandUtil) SetHands() {

	p.hands = make([]interface{}, 10, 20)

	p.hands[models.HandTypeRoyalStraightFlush] = p.CheckRoyalStraightFlush
	p.hands[models.HandTypeStraightFlush] = p.CheckStraightFlush
	p.hands[models.HandTypePoker] = p.CheckPoker
	p.hands[models.HandTypeFullHouse] = p.CheckFullHouse
	p.hands[models.HandTypeFlush] = p.CheckFlush
	p.hands[models.HandTypeStrait] = p.CheckStraight
	p.hands[models.HandTypeTriple] = p.CheckTriple
	p.hands[models.HandTypeTwoPair] = p.CheckTwoPairs
	p.hands[models.HandTypeOnePair] = p.CheckOnePair
	p.hands[models.HandTypeTitle] = p.CheckTitle
}

// CheckRoyalStraightFlush 로얄스트레이트 플러쉬
func (p *PokerHandUtil) CheckRoyalStraightFlush(cards []int) models.HandResult {

	result := models.HandResult{}
	handsCount := 4
	cardsCount := 5
	hands := [4][5]int{
		{0, 9, 10, 11, 12},
		{13, 22, 23, 24, 25},
		{26, 35, 36, 37, 38},
		{39, 48, 49, 50, 51}}

	matchCount := 0
	compare1 := -1
	compare2 := -1

	for i := 0; i < handsCount; i++ {
		matchCount = 0
		result.InitializeMadeCard()
		for k := 0; k < cardsCount; k++ {
			for j := 0; j < len(cards); j++ {
				compare1 = hands[i][k]
				compare2 = cards[j]

				if compare2 == -1 {
					continue
				}

				if compare1 == compare2 {
					matchCount++
					result.AddMadeCard(compare1)
				}
			}
		}
		if matchCount == cardsCount {
			result.HandType = models.HandTypeRoyalStraightFlush
			result.CardType = i + 1

			break
		}
		matchCount = 0
	}

	return result
}

// CheckStraightFlush 스트레이트 플러쉬
func (p *PokerHandUtil) CheckStraightFlush(cards []int) models.HandResult {
	result := models.HandResult{}
	handsCount := 36
	cardsCount := 5
	hands := [36][5]int{
		{0, 1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
		{4, 5, 6, 7, 8},
		{5, 6, 7, 8, 9},
		{6, 7, 8, 9, 10},
		{7, 8, 9, 10, 11},
		{8, 9, 10, 11, 12},

		{13, 14, 15, 16, 17},
		{14, 15, 16, 17, 18},
		{15, 16, 17, 18, 19},
		{16, 17, 18, 19, 20},
		{17, 18, 19, 20, 21},
		{18, 19, 20, 21, 22},
		{19, 20, 21, 22, 23},
		{20, 21, 22, 23, 24},
		{21, 22, 23, 24, 25},

		{26, 27, 28, 29, 30},
		{27, 28, 29, 30, 31},
		{28, 29, 30, 31, 32},
		{29, 30, 31, 32, 33},
		{30, 31, 32, 33, 34},
		{31, 32, 33, 34, 35},
		{32, 33, 34, 35, 36},
		{33, 34, 35, 36, 37},
		{34, 35, 36, 37, 38},

		{39, 40, 41, 42, 43},
		{40, 41, 42, 43, 44},
		{41, 42, 43, 44, 45},
		{42, 43, 44, 45, 46},
		{43, 44, 45, 46, 47},
		{44, 45, 46, 47, 48},
		{45, 46, 47, 48, 49},
		{46, 47, 48, 49, 50},
		{47, 48, 49, 50, 51}}
	matchCount := 0
	compare1 := -1
	compare2 := -1

	for i := 0; i < handsCount; i++ {
		matchCount = 0
		result.InitializeMadeCard()

		for k := 0; k < cardsCount; k++ {
			for j := 0; j < len(cards); j++ {
				compare1 = hands[i][k]
				compare2 = cards[j]

				if compare2 == -1 {
					continue
				}

				if compare1 == compare2 {
					matchCount++
					result.AddMadeCard(compare1)
				}
			}
		}

		if matchCount == cardsCount {
			result.HandType = models.HandTypeStraightFlush
			result.CardType = ((i-(i%9))/9 + 1)
			result.Hands[0] = 4 + (i % 9)

			break
		}

		matchCount = 0
	}

	return result
}

// CheckPoker 포카드
func (p *PokerHandUtil) CheckPoker(cards []int) models.HandResult {
	result := models.HandResult{}
	matchCount := 0
	matchIndex := 0
	i := 0
	for i = 0; i < len(cards); i++ {
		matchCount = 0
		for j := 0; j < len(cards); j++ {
			if cards[j] == -1 {
				continue
			}
			if cards[i]%13 == cards[j]%13 {
				if matchCount == 3 {
					matchIndex = cards[i] % 13
					result.AddMadeCard(cards[i])
				}
				matchCount++
			}
		}

		if matchCount == 4 {
			result.HandType = models.HandTypePoker
			if matchIndex == 0 {
				matchIndex = 13
			} else {
				result.Hands[0] = matchIndex
			}
		}
	}
	return result
}

// CheckFullHouse 풀하우스
func (p *PokerHandUtil) CheckFullHouse(cards []int) models.HandResult {
	result := models.HandResult{}
	i := 0
	j := 0
	hasTriple := false
	matchCount := 0
	matchIndex1 := 0
	//matchIndex2 := 0
	isFullHouse := false
	for i = 0; i < len(cards); i++ {
		matchCount = 0
		for j = 0; j < len(cards); j++ {
			if cards[j] == -1 {
				continue
			}
			if cards[i]%13 == cards[j]%13 {
				if matchCount == 2 {
					matchIndex1 = cards[i] % 13
					result.AddMadeCard(cards[i])
				}
				matchCount++
			}
		}

		if matchCount >= 3 {
			hasTriple = true
		}
	}
	if hasTriple == true {
		for i = 0; i < len(cards); i++ {
			matchCount = 0

			for j = 0; j < len(cards); j++ {
				if cards[j] == -1 {
					continue
				}
				if cards[i]%13 == cards[j]%13 && cards[i]%13 != matchIndex1 {
					if matchCount == 1 {
						//matchIndex2 = cards[i] % 13
						result.AddMadeCard(cards[i])
					}
					matchCount++
				}
			}

			if matchCount >= 2 {
				isFullHouse = true
				result.HandType = models.HandTypeFullHouse
			}
		}
	}
	if isFullHouse == true {
		hand1 := -100
		hand2 := -100

		for i = 0; i < len(cards); i++ {
			matchCount = 0
			matchIndex1 = -1

			for j = 0; j < len(cards); j++ {
				if cards[j] == -1 {
					continue
				}

				if cards[i]%13 == cards[j]%13 {
					if matchCount == 1 {
						matchIndex1 = cards[i] % 13
						if matchIndex1 == 0 {
							matchIndex1 = 13
						}
					}
					matchCount++
				}

				if matchIndex1 == -1 {
					continue
				}

				if matchCount == 2 {
					// 페어일떄는 가장큰 숫자로 대체
					if matchIndex1 > hand2 {
						hand2 = matchIndex1
					}
				} else if matchCount == 3 {
					// 트리플일때는 트리플 값을 세팅
					if matchIndex1 > hand1 {
						hand1 = matchIndex1
					} else {
						if matchIndex1 > hand2 {
							hand2 = matchIndex1
						}
					}
				}
			}
		}
		result.Hands[0] = hand1
		result.Hands[1] = hand2
	}

	return result
}

// CheckFlush 풀러쉬
func (p *PokerHandUtil) CheckFlush(cards []int) models.HandResult {
	result := models.HandResult{}
	i := 0
	j := 0
	matchCount := 0
	count := 0
	matchType := models.CardTypeNone

	for i = 0; i < len(cards); i++ {
		matchCount = 0
		for j = 0; j < len(cards); j++ {
			if cards[j] == -1 {
				continue
			}
			if GetCardType(cards[i]) == GetCardType(cards[j]) {
				if matchCount == 4 {
					matchType = GetCardType(cards[i])
					result.AddMadeCard(cards[i])
					result.AddMadeCard(cards[j])
				}
				matchCount++
			}
		}

		if matchCount >= 5 {
			result.HandType = models.HandTypeFlush
			result.CardType = matchType
			count = 0

			for j = 0; j < len(cards); j++ {
				if GetCardType(cards[j]) == matchType {
					result.Hands[count] = cards[j] % 13
					if result.Hands[count] == 0 {
						result.Hands[count] = 13
					}

					count++
				}
			}

			// 정렬
			slice.Sort(result.Hands[:], func(i, j int) bool {
				return result.Hands[i] < result.Hands[j]
			})
			// reverse
			for i := len(result.Hands)/2 - 1; i >= 0; i-- {
				opp := len(result.Hands) - 1 - i
				result.Hands[i], result.Hands[opp] = result.Hands[opp], result.Hands[i]
			}
		}
	}

	return result
}

// CheckStraight 스트레이트
func (p *PokerHandUtil) CheckStraight(cards []int) models.HandResult {
	result := models.HandResult{}
	hands := [10]string{
		"0,1,2,3,4",
		"1,2,3,4,5",
		"2,3,4,5,6",
		"3,4,5,6,7",
		"4,5,6,7,8",
		"5,6,7,8,9",
		"6,7,8,9,10",
		"7,8,9,10,11",
		"8,9,10,11,12",
		"9,10,11,12"}
	i, j, k := 0, 0, 0

	compareCards := make([]int, len(cards), len(cards))

	for i = 0; i < len(cards); i++ {
		compareCards[i] = cards[i] % 13
	}
	cardStr := GetCardStr(compareCards)

	for i = 0; i < len(hands); i++ {
		result.InitializeMadeCard()
		if i == 9 {
			if strings.Index(cardStr, hands[i]) >= 0 && GetIntArrayIndexOf(compareCards, 0) >= 0 {
				result.HandType = models.HandTypeStrait
				result.Hands[0] = 13

				selectedHands := strings.Split(hands[i], ",")

				for j = 0; j < len(compareCards); j++ {
					for k = 0; k < len(selectedHands); k++ {
						parseInt, _ := strconv.Atoi(selectedHands[k])
						if compareCards[j] == parseInt {
							result.AddMadeCard(cards[j])
							break
						}
					}

					if cards[j]%13 == 0 {
						result.AddMadeCard(cards[j])
					}
				}
				break
			}
		} else {
			if strings.Index(cardStr, hands[i]) >= 0 {
				result.HandType = models.HandTypeStrait
				result.Hands[0] = 13 - (9 - i)

				selectedHands := strings.Split(hands[i], ",")

				for j = 0; j < len(compareCards); j++ {
					for k = 0; k < len(selectedHands); k++ {
						parseInt, _ := strconv.Atoi(selectedHands[k])
						if compareCards[j] == parseInt {
							result.AddMadeCard(cards[j])
							break
						}
					}
				}
				break
			}
		}
	}

	return result
}

// CheckTriple 트리플
func (p *PokerHandUtil) CheckTriple(cards []int) models.HandResult {
	result := models.HandResult{}
	i := 0
	j := 0
	matchCount := 0
	matchIndex := -1
	kickCount := 0

	for i = 0; i < len(cards); i++ {
		matchCount = 0
		if cards[i] == -1 {
			continue
		}

		for j = 0; j < len(cards); j++ {
			if cards[i]%13 == cards[j]%13 {
				if matchCount == 2 {
					matchIndex = cards[i] % 13
					if matchIndex == 0 {
						matchIndex = 13
					}

					result.Hands[0] = matchIndex

					result.AddMadeCard(cards[i])
					result.AddMadeCard(cards[j])
				}
				matchCount++
			}
		}

		if matchCount == 3 {
			result.HandType = models.HandTypeTriple
		}
	}
	var card int
	for i = 0; i < len(cards); i++ {
		card = cards[i] % 13
		if card == 0 {
			card = 13
		}

		if card != matchCount {
			result.Kicks[kickCount] = card
			kickCount++
		}
	}

	// 정렬
	slice.Sort(result.Kicks[:], func(i, j int) bool {
		return result.Kicks[i] < result.Kicks[j]
	})

	return result
}

// CheckTwoPairs 투페어
func (p *PokerHandUtil) CheckTwoPairs(cards []int) models.HandResult {
	result := models.HandResult{}
	i := 0
	j := 0
	k := 0
	matchCount := 0
	matchIndex := -1
	endCount := 0
	exist := false
	kickCount := 0

	for i = 0; i < len(cards); i++ {
		matchCount = 0
		if cards[i] == -1 {
			continue
		}
		for j = 0; j < len(cards); j++ {
			if cards[i]%13 == cards[j]%13 {
				if matchCount == 1 {
					matchIndex = cards[i] % 13
					if matchIndex == 0 {
						matchIndex = 13
					}
					exist = false

					for k = 0; k < endCount; k++ {
						if result.Hands[k] == matchIndex {
							exist = true
						}
					}

					if exist == false {
						result.Hands[endCount] = matchIndex
						matchCount++

						result.AddMadeCard(cards[i])
						result.AddMadeCard(cards[j])
					}
				} else {
					matchCount++
				}
			}
		}

		if matchCount == 2 {
			matchCount = 0
			endCount++
		}
	}

	if endCount >= 2 {
		result.HandType = models.HandTypeTwoPair

		// 정렬
		slice.Sort(result.Hands[:], func(i, j int) bool {
			return result.Hands[i] < result.Hands[j]
		})
		// reverse
		for i := len(result.Hands)/2 - 1; i >= 0; i-- {
			opp := len(result.Hands) - 1 - i
			result.Hands[i], result.Hands[opp] = result.Hands[opp], result.Hands[i]
		}
	}

	var card int
	for i = 0; i < len(cards); i++ {
		card = cards[i] % 13
		if card == 0 {
			card = 13
		}

		if card != result.Hands[0] && card != result.Hands[1] {
			result.Kicks[kickCount] = card
			kickCount++
		}
	}
	return result
}

// CheckOnePair 원페어
func (p *PokerHandUtil) CheckOnePair(cards []int) models.HandResult {
	result := models.HandResult{}
	i := 0
	j := 0
	k := 0
	matchCount := 0
	kickCount := 0

	for i = 0; i < len(cards); i++ {
		matchCount = 0
		if cards[i] == -1 {
			continue
		}
		for j = 0; j < len(cards); j++ {
			if cards[i]%13 == cards[j]%13 {
				if matchCount == 1 {
					result.Hands[0] = cards[i] % 13
					if result.Hands[0] == 0 {
						result.Hands[0] = 13
					}

					for k = 0; k < len(cards); k++ {
						index := GetCardOrder(cards, k)
						temp := index % 13
						if temp == 0 {
							temp = 13
						}
						if temp == result.Hands[0] {
							result.AddMadeCard(cards[i])
							result.AddMadeCard(cards[j])

							result.CardType = GetCardType(index)
							break
						}
					}
				}
				matchCount++
			}
		}

		if matchCount >= 2 {
			result.HandType = models.HandTypeOnePair
		}
	}

	var card int
	for i = 0; i < len(cards); i++ {
		card = cards[i] % 13
		if card == 0 {
			card = 13
		}

		if card != result.Hands[0] {
			result.Kicks[kickCount] = card
			kickCount++
		}
	}

	// 정렬
	slice.Sort(result.Kicks[:], func(i, j int) bool {
		return result.Kicks[i] < result.Kicks[j]
	})

	return result
}

// CheckTitle 타이틀
func (p *PokerHandUtil) CheckTitle(cards []int) models.HandResult {
	result := models.HandResult{}
	result.HandType = models.HandTypeTitle

	for i := 0; i < len(cards); i++ {
		result.Hands[i] = GetCardOrder(cards, i) % 13
		if result.Hands[i] == 0 {
			result.Hands[i] = 13
		}
		result.AddMadeCard(cards[i])
	}
	return result
}

//GetCardType 카드인덱스로 카드타입을 반환
func GetCardType(cardIndex int) int {
	cardType := models.CardTypeNone

	if cardIndex >= 0 && cardIndex <= 12 {
		cardType = models.CardTypeSpade //스페이드
	}
	if cardIndex >= 13 && cardIndex <= 25 {
		cardType = models.CardTypeDiamond //다이아몬드
	}
	if cardIndex >= 26 && cardIndex <= 38 {
		cardType = models.CardTypeHeart //하트
	}
	if cardIndex >= 39 && cardIndex <= 51 {
		cardType = models.CardTypeClover //클로버
	}
	return cardType
}

//GetCardStr 카드인덱스모음을 문자열로 변환
func GetCardStr(cards []int) string {
	str := ""
	tempCards := make([]int, len(cards))
	count := 0
	copy(tempCards, cards)
	// 정렬
	slice.Sort(tempCards[:], func(i, j int) bool {
		return tempCards[i] < tempCards[j]
	})

	for i := 0; i < len(tempCards); i++ {
		if tempCards[i] == -1 {
			continue
		}
		if count > 0 {
			str += ","
		}
		str += strconv.Itoa(tempCards[i])
		count++
	}
	return str
}

// GetCardOrder 카드정렬을 가져온다.
func GetCardOrder(cards []int, orderNo int) int {
	i := 0
	value := 0
	order := make([]int, len(cards))

	for i = 0; i < len(cards); i++ {
		if cards[i]%13 == 0 {
			value = 13
		} else {
			value = cards[i] % 13
		}
		value = value*4 - (GetCardType(cards[i]) - 1)

		order[i] = value
	}

	// 정렬
	slice.Sort(order[:], func(i, j int) bool {
		return order[i] < order[j]
	})

	ret := order[orderNo]
	cardType := ret % 4
	if cardType == 0 {
		cardType = 1
	} else if cardType == 1 {
		cardType = 2
	} else if cardType == 2 {
		cardType = 3
	} else if cardType == 3 {
		cardType = 4
	}

	ret = (ret+cardType-1)/4 + 13*(cardType-1)

	return ret
}
