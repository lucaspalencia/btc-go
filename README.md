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

## Development

To run `btc-go` on your local computer, following this step-by-step instruction:

```
$ git clone git@github.com:lucaspalencia/btc-go.git
$ cd btc-go
$ go run main.go price
```