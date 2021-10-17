package constvar

import "github.com/shopspring/decimal"

// Decimal is a github.com/shopspring/decimal.Decimal.
func Decimal(c decimal.Decimal) func() decimal.Decimal { return func() decimal.Decimal { return c } }
