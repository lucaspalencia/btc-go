# BTC-GO

See bitcoin market data in USD, EUR and BRL

## Usage

`btc-go` provides these commands.

- price, p

### Price

Get current bitcoin price for the chosen currency.

#### Options

##### `-c <currency>` or `--currency <currency>`

Choose a currency: usd, eur or brl

If you don't specify the option, the currency by default is usd.

```
$ btc-go price
$ btc-go price -c eur
$ btc-go price -c brl
```

![btc-go price](https://user-images.githubusercontent.com/7226038/114094407-a670f580-9892-11eb-95f5-b9da6deca24f.gif)

## Development

To run `btc-go` on your local computer, following this step-by-step instruction:

```
$ git clone git@github.com:lucaspalencia/btc-go.git
$ cd btc-go
$ go run main.go price
```