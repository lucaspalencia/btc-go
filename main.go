package main

import (
	"fmt"
	"os"

	"github.com/lucaspalencia/btc-go/app"
)

func main() {
	app := app.Build()

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("")
		fmt.Println("Please run btc-go again.\nIf it still does not work, feel free to open an issue on https://github.com/lucaspalencia/btc-go/issues")
	}
}
