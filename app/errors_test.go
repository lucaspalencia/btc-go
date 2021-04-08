package app

import "testing"

func TestErrors(t *testing.T) {
	commandName := "show_price_command"
	text := "invalid currency"

	error := NewAppError(commandName, text)

	expected := "Oops, " + commandName + " command goes wrong: " + text
	result := error.Error()

	if expected != result {
		t.Errorf("Expected %s but receives %s", expected, result)
	}
}
