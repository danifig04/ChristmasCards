package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/danifig04/ChristmasCards/christmascards"
)

//will need to delete storage.go
const filename = "christmascards.json"

func Load() error {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var savedCards []*christmascards.Card
	err = json.Unmarshal(fileContents, &savedCards)
	if err != nil {
		return err
	}

	christmascards.SetCards(savedCards)

	return nil
}

func Save() error {
	christmasCardsList := christmascards.ListCard()

	christmasCardsListBytes, err := json.MarshalIndent(christmasCardsList, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, christmasCardsListBytes, 0775)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, christmasCardsListBytes, 0775)
	if err != nil {
		return err
	}

	return nil
}
