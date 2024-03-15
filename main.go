package main

import (
	"github.com/Noofbiz/go-card/smartcard"
)

func main() {
	ctx, err := smartcard.EstablishContext(smartcard.SCOPE_USER)
	if err != nil {
		panic("Unable to establish context with card reader. Error was: [[ " + err.Error() + " ]]")
	}
	defer ctx.Release()

	readers := ctx.ListReaders()
	if len(readers) == 0 {
		//no card detected
	}
}
