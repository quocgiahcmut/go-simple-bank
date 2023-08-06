package util

// USD VND CAD EUR
const (
	USD = "USD"
	VND = "VND"
	CAD = "CAD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, VND, CAD, EUR:
		return true
	}
	return false
}
