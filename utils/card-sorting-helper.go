package utils

import (
	"holdempoker/models"
	"math/rand"
)

//CardSortingHelper 카드소팅핼퍼
type CardSortingHelper struct {
	cards []int
}

//SetCards 카드를 세팅한다.
func (c *CardSortingHelper) SetCards() {
	c.cards = make([]int, models.CardCount, models.CardCount)
	for i := 0; i < models.CardCount; i++ {
		c.cards[i] = i
	}
}

//Shuffle 카드를 섞는다.
func (c *CardSortingHelper) Shuffle() {
	rand.Shuffle(len(c.cards), func(i, j int) {
		c.cards[i], c.cards[j] = c.cards[j], c.cards[i]
	})
}

//Pop 카드를 한장씩 뺀다
func (c *CardSortingHelper) Pop() int {
	p := -1
	if len(c.cards) > 0 {
		p, c.cards = c.cards[0], c.cards[1:]
	}
	return p
}

//Initialize 카드 초기화
func (c *CardSortingHelper) Initialize() {
	c.SetCards()
	c.Shuffle()
}
