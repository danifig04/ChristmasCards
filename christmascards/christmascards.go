package christmascards

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	insertCardsQuery = "INSERT INTO cards (FamilyNames, Address, PostagePrice, MembersNames) VALUES (?, ?, ?, ?)"
)

var (
	cards []*Card
)

type CardsService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *CardsService {
	return &CardsService{
		db: db,
	}
}

type Card struct {
	FamilyName   string
	Address      string
	PostagePrice float64
	MembersNames []string
	IsSent       bool
}

//TODO: Add MembersNames

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

func (c *CardsService) SaveCard() error {
	christmasCardsList := ListCard()

	for _, card := range christmasCardsList {
		famMembers := strings.Join(card.MembersNames, ",")
		_, err := c.db.Exec(insertCardsQuery, card.FamilyName, card.Address, card.PostagePrice, famMembers)
		if err != nil {
			return err
		}
	}

	return nil
}
