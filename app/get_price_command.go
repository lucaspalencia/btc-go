package app

import (
	"btc-go/config"
	"btc-go/services"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func getPriceCommand(c *cli.Context) error {
	currency := strings.ToLower(c.String("currency"))

	if !isValidCurrency(currency) {
		return errors.New("Invalid currency")
	}

	data, err := services.GetBitcoinData(currency)

	if err != nil {
		return err
	}

	p := message.NewPrinter(language.English)
	p.Printf("Current bitcoin price in "+strings.ToUpper(currency)+": "+getCurrencySymbol(currency)+"%d\n", data.Price)

	unixTimeUTC := time.Unix(int64(data.UpdatedAt), 0)

	fmt.Println("Updated at:", unixTimeUTC.Format(config.New().DATETIME_FORMAT))

	return nil
}
