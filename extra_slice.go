package constvar

import "github.com/shopspring/decimal"

// Decimals is a github.com/shopspring/decimal.Decimal.
func Decimals(c []decimal.Decimal) func() []decimal.Decimal {
	return func() []decimal.Decimal { return c }
}
