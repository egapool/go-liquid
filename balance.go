package liquid

import "github.com/shopspring/decimal"

type Balance struct {
	Currency string          `json:"currency"`
	Balance  decimal.Decimal `json:"balance"`
}
