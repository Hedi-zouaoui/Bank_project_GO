package util

// Constants for all supported currencies
const (
	USD = "usd"
	EUR = "euro"
	TND = "TND"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TND:
		return true
	}
	return false
}