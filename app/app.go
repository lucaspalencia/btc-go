package app

import (
	"github.com/urfave/cli/v2"
)

func Build() *cli.App {
	app := cli.NewApp()
	app.Name = "Bitcoin Go"
	app.Usage = "Get bitcoin market data and convert to USD/EUR/BRL"

	app.Commands = buildCommands()

	return app
}

func buildCommands() []*cli.Command {
	currencyFlag := &cli.StringFlag{
		Name:    "currency",
		Value:   "usd",
		Aliases: []string{"c"},
	}

	return []*cli.Command{
		{
			Name:    "price",
			Usage:   "Get current bitcoin price for the chosen currency",
			Aliases: []string{"p"},
			Flags:   []cli.Flag{currencyFlag},
			Action:  getPriceCommand,
		},
		{
			Name:   "market",
			Usage:  "Get the current bitcoin market data",
			Flags:  []cli.Flag{currencyFlag},
			Action: getMarketDataCommand,
		},
	}
}
