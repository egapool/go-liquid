package liquid

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID                  string          `json:"id"`
	ProductType         string          `json:"product_type"`
	Code                string          `json:"code"`
	Name                string          `json:"name"`
	MarketAsk           decimal.Decimal `json:"market_ask"`
	MarketBid           decimal.Decimal `json:"market_bid"`
	Indicator           int             `json:"indicator"`
	Currency            string          `json:"currency"`
	CurrencyPairCode    string          `json:"currency_pair_code"`
	Symbol              string          `json:"symbol"`
	BtcMinimumWithdraw  interface{}     `json:"btc_minimum_withdraw"`
	FiatMinimumWithdraw interface{}     `json:"fiat_minimum_withdraw"`
	PusherChannel       string          `json:"pusher_channel"`
	TakerFee            json.Number     `json:"taker_fee"`
	MakerFee            json.Number     `json:"maker_fee"`
	LowMarketBid        json.Number     `json:"low_market_bid"`
	HighMarketAsk       json.Number     `json:"high_market_ask"`
	Volume24H           json.Number     `json:"volume_24h"`
	LastPrice24H        json.Number     `json:"last_price_24h"`
	LastTradedPrice     decimal.Decimal `json:"last_traded_price"`
	LastTradedQuantity  json.Number     `json:"last_traded_quantity"`
	QuotedCurrency      string          `json:"quoted_currency"`
	BaseCurrency        string          `json:"base_currency"`
	Disabled            bool            `json:"disabled"`
	MarginEnabled       bool            `json:"margin_enabled"`
	CfdEnabled          bool            `json:"cfd_enabled"`
	LastEventTimestamp  string          `json:"last_event_timestamp"`
	ExchangeRate        int             `json:"exchange_rate"`
}
