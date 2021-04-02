package main

import (
	"btc-go/app"
	"log"
	"os"
)

func main() {
	app := app.Build()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
