package liquid

import "encoding/json"

type TradeResponse struct {
	Models      []Trade `json:"models"`
	CurrentPage int     `json:"current_page"`
	TotalPages  int     `json:"total_pages"`
}

type Trade struct {
	ID               int         `json:"id"`
	CurrencyPairCode string      `json:"currency_pair_code"`
	Status           string      `json:"status"`
	Side             string      `json:"side"`
	MarginUsed       string      `json:"margin_used"`
	OpenQuantity     string      `json:"open_quantity"`
	CloseQuantity    string      `json:"close_quantity"`
	Quantity         string      `json:"quantity"`
	LeverageLevel    int         `json:"leverage_level"`
	ProductCode      string      `json:"product_code"`
	ProductID        int         `json:"product_id"`
	OpenPrice        float64     `json:"open_price,string"`
	ClosePrice       float64     `json:"close_price,string"`
	TraderID         int         `json:"trader_id"`
	OpenPnl          json.Number `json:"open_pnl"`
	ClosePnl         json.Number `json:"close_pnl"`
	Pnl              json.Number `json:"pnl"`
	StopLoss         json.Number `json:"stop_loss"`
	TakeProfit       string      `json:"take_profit"`
	FundingCurrency  string      `json:"funding_currency"`
	CreatedAt        int         `json:"created_at"`
	UpdatedAt        int         `json:"updated_at"`
	TotalInterest    string      `json:"total_interest"`
}
