package app

import (
	"github.com/lucaspalencia/btc-go/config"
	"github.com/urfave/cli/v2"
)

func Build() *cli.App {
	config := config.New()

	app := cli.NewApp()
	app.Name = config.APP_NAME
	app.Usage = config.APP_USAGE
	app.Version = config.APP_VERSION

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
			Action:  showPriceCommand,
		},
	}
}
