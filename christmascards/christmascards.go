package christmascards

import "fmt"

var (
	cards []*Card
)

type Card struct {
	FamilyName   string
	Address      string
	PostagePrice float64
	MembersNames []string
	IsSent       bool
}

func AddCard(card *Card) {
	cards = append(cards, card)
}

func SetCards(c []*Card) {
	cards = c
}

func ListCard() []*Card {
	return cards
}

func (card *Card) CheckOffCard() {
	card.IsSent = true
}

func (card *Card) AddMembersName(name string) {
	card.MembersNames = append(card.MembersNames, name)
}

func CalculateTotalPostage() float64 {
	totalPostage := 0.0
	for _, value := range cards {
		totalPostage = totalPostage + value.PostagePrice
		fmt.Println(totalPostage)
	}
	return totalPostage
}
