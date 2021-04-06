package app

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/leekchan/accounting"
	"github.com/lucaspalencia/btc-go/config"
	"github.com/lucaspalencia/btc-go/services"
	"github.com/urfave/cli/v2"
)

func showPriceCommand(c *cli.Context) error {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Loading bitcoin price data"

	currency := strings.ToLower(c.String("currency"))

	s.Start()

	if !isValidCurrency(currency) {
		return NewAppError(c.Command.Name, "Invalid currency")
	}

	bitcoinData, err := services.GetBitcoinData(currency)

	if err != nil {
		return NewAppError(c.Command.Name, "Service error")
	}

	fmt.Println("")

	printPrice(currency, bitcoinData.Price)
	printPriceChange(bitcoinData.MarketPercentage)
	printUpdatedAt(bitcoinData.UpdatedAt)

	s.Stop()

	return nil
}

func printPrice(currency string, price int) {
	colorFunc := color.New(color.FgCyan).SprintFunc()
	lc := accounting.LocaleInfo[strings.ToUpper(currency)]
	ac := accounting.Accounting{Symbol: lc.ComSymbol, Thousand: lc.ThouSep}

	fmt.Println("Current bitcoin price in "+strings.ToUpper(currency)+":", colorFunc(ac.FormatMoney(price)))
}

func printPriceChange(changePercentage float64) {
	colorFunc := color.New(getColor(changePercentage)).SprintFunc()
	var priceChange = "" + strconv.FormatFloat(changePercentage, 'f', 2, 64) + "%"

	if changePercentage > 0 {
		priceChange = "+" + priceChange
	}

	fmt.Println("Change in the last 24 hours:", colorFunc(priceChange))
}

func printUpdatedAt(date int) {
	colorFunc := color.New(color.FgYellow).SprintFunc()
	unixTimeUTC := time.Unix(int64(date), 0)
	dateFormatted := unixTimeUTC.Format(config.New().DATETIME_FORMAT)

	fmt.Println("Updated at:", colorFunc(dateFormatted))
}

func getColor(changePercentage float64) color.Attribute {
	if changePercentage < 0 {
		return color.FgRed
	}

	return color.FgGreen
}
