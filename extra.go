package constvar

import "github.com/shopspring/decimal"

// Decimal is a github.com/shopspring/decimal.Decimal.
type Decimal decimal.Decimal

func (c Decimal) Val() decimal.Decimal { return decimal.Decimal(c) }
