package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/danifig04/ChristmasCards/christmascards"
	"github.com/danifig04/ChristmasCards/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

const (
	addCard          = "Add Card"
	numberOfCards    = "Calculate number of cards needed"
	calculatePostage = "Calculate Total Postage"
	showListOfCards  = "List of cards"
	done             = "Done"
)

func main() {
	db, err := db.ConnectDatabase("db/christmascards_db.config")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	cardService := christmascards.NewService(db)

	for {
		fmt.Println()
		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{
				addCard,
				numberOfCards,
				calculatePostage,
				showListOfCards,
				done,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		finished := false
		switch result {
		case addCard:
			err := addCardPrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		case numberOfCards:
			PrintNumberofCards()

		case calculatePostage:
			PrintCalculatePostage()

		case showListOfCards:
			PrintShowListOfCards()

		case done:
			finished = true
		}

		if finished {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	//save to sql
	cardService.SaveCard()

	print("Done with cards")
	// no longer saves to json
	//err = storage.Save()
	// if err != nil {
	// 	fmt.Println("Error saving file", err)
	// 	return
	// }
}
func addCardPrompt() error {

	namePrompt := promptui.Prompt{
		Label: "Family name",
	}
	name, err := namePrompt.Run()
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(name)

	addressPrompt := promptui.Prompt{
		Label: "Family address",
	}
	address, err := addressPrompt.Run()
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(address)

	postagePricePrompt := promptui.Prompt{
		Label: "Postage cost",
	}

	inputStr, err := postagePricePrompt.Run()
	if err != nil {
		return err
	}
	postagePrice, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(postagePrice)

	newCard := &christmascards.Card{
		FamilyName:   name,
		Address:      address,
		PostagePrice: postagePrice,
	}

	christmascards.AddCard(newCard)

	fmt.Println("Added card", newCard)

	return nil
}

func Save() error {
	christmasCardsList := christmascards.ListCard()

	christmasCardsListBytes, err := json.MarshalIndent(christmasCardsList, "", "    ")
	if err != nil {
		return err
	}

	filename := "christmascards.json"
	err = ioutil.WriteFile(filename, christmasCardsListBytes, 0775)
	if err != nil {
		return err
	}

	return nil
}

func PrintCalculatePostage() {
	calculatePostage := christmascards.CalculateTotalPostage()

	if calculatePostage == 0 {
		fmt.Println("No postage cost")
		return
	}
	fmt.Println("Total postage = $", calculatePostage)
}

func PrintNumberofCards() {
	numberOfCards := len(christmascards.ListCard())
	fmt.Println("You need", numberOfCards, "cards")
}

func cardsLeftToSendPrompt() error {
	availableCard := christmascards.ListCard()

	if len(availableCard) == 0 {
		fmt.Println("No cards left to send!")
		return nil
	}

	var options []string
	for _, card := range availableCard {
		options = append(options, card.FamilyName)
	}
	return nil
}
func PrintShowListOfCards() {
	namesOnList := christmascards.ListCard()

	if len(namesOnList) == 0 {
		fmt.Println("No cards to send!")
		return
	}
	for _, v := range namesOnList {
		fmt.Println(v.FamilyName)
	}
}
