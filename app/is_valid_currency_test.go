package app

import "testing"

func TestIsValidCurrency(t *testing.T) {
	testAssert := func(t testing.TB, expected, result bool) {
		t.Helper()
		if expected != result {
			t.Errorf("Expected %t but receives %t", expected, result)
		}
	}

	t.Run("Invalid currency", func(t *testing.T) {
		currency := "cny"
		expected := false

		result := isValidCurrency(currency)

		testAssert(t, expected, result)
	})

	t.Run("USD currency", func(t *testing.T) {
		currency := "usd"
		expected := true

		result := isValidCurrency(currency)

		testAssert(t, expected, result)
	})

	t.Run("EUR currency", func(t *testing.T) {
		currency := "eur"
		expected := true

		result := isValidCurrency(currency)

		testAssert(t, expected, result)
	})

	t.Run("BRL currency", func(t *testing.T) {
		currency := "brl"
		expected := true

		result := isValidCurrency(currency)

		testAssert(t, expected, result)
	})
}
