package constvar

import "github.com/shopspring/decimal"

// Decimals is a github.com/shopspring/decimal.Decimal.
type Decimals []decimal.Decimal

func (c Decimals) Val() []decimal.Decimal { return []decimal.Decimal(c) }
