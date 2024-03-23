package main

import (
	"fmt"

	"github.com/Noofbiz/go-card/smartcard"
)

func main() {
	ctx, err := smartcard.EstablishContext(smartcard.SCOPE_USER)
	if err != nil {
		panic("Unable to establish context with card reader. Error was: [[ " + err.Error() + " ]]")
	}
	defer ctx.Release()

	readers, err := ctx.ListReaders()
	if len(readers) == 0 {
		//no card detected
	}
	if err != nil {
		panic("Unable to find card readers. Error was: [[ " + err.Error() + " ]]")
	}

	for _, reader := range readers {
		fmt.Printf("Reader found.\nName is [[ %v ]].\n", reader.Name())
		if !reader.IsCardPresent() {
			fmt.Println("No card found in reader. Waiting for card.")
			ctx.WaitForCardPresent()
		}
		fmt.Println("Card detected. Attempting to connect to card.")
		card, err := reader.Connect()
		if err != nil {
			panic("Unable to get card from reader even though it's inserted. Error was: [[ " + err.Error() + " ]]")
		}
		fmt.Println("Card connection established.")
		fmt.Printf("Card ATR: %s\n", card.ATR())
	}
}
