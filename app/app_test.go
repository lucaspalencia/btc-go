package app

import (
	"reflect"
	"testing"

	"github.com/lucaspalencia/btc-go/config"
	"github.com/urfave/cli/v2"
)

func TestApp(t *testing.T) {
	config := config.New()
	app := Build()

	t.Run("App type", func(t *testing.T) {
		expected := reflect.TypeOf(&cli.App{})
		result := reflect.TypeOf(app)

		if expected != result {
			t.Errorf("Expected %s but receives %s", expected, result)
		}
	})

	t.Run("App name", func(t *testing.T) {
		expected := config.APP_NAME
		result := app.Name

		if expected != result {
			t.Errorf("Expected %s but receives %s", expected, result)
		}
	})

	t.Run("App usage", func(t *testing.T) {
		expected := config.APP_USAGE
		result := app.Usage

		if expected != result {
			t.Errorf("Expected %s but receives %s", expected, result)
		}
	})

	t.Run("App version", func(t *testing.T) {
		expected := config.APP_VERSION
		result := app.Version

		if expected != result {
			t.Errorf("Expected %s but receives %s", expected, result)
		}
	})
}
