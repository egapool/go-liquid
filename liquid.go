package liquid

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	API_BASE    = "https://api.liquid.com" // Liquid API endpoint
	API_VERSION = "2"
)

// New returns an instantiated liquid struct
func New(apiKey, apiSecret string) *Liquid {
	client := NewClient(apiKey, apiSecret)
	return &Liquid{client}
}

// Liquid represent a liquid client
type Liquid struct {
	client *client
}

// SetDebug is set enable/disable http request/response dump
func (l *Liquid) SetDebug(enable bool) {
	l.client.debug = enable
}

// GetTicker is used to get the current ticker values for a market.
func (l *Liquid) GetTicker(productID int) (ticker Ticker, err error) {
	r, err := l.client.do("GET", "/products/"+strconv.Itoa(productID), "", false)
	if err != nil {
		return
	}
	var product Product
	if err = json.Unmarshal(r, &product); err != nil {
		return
	}
	ticker = Ticker{
		product.MarketBid,
		product.MarketAsk,
		product.LastTradedPrice,
	}
	return
}

// GetOrderBook is
func (l *Liquid) GetOrderBook(productID int, full bool) (orderBook OrderBook, err error) {
	query := "/products/" + strconv.Itoa(productID) + "/price_levels"
	if full {
		query += "?full=1"
	}
	r, err := l.client.do("GET", query, "", false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &orderBook); err != nil {
		return
	}
	return

}

// GetAllAccountBalances is used to get all balances that you have.
func (l *Liquid) GetAllAccountBalances() (balances []Balance, err error) {
	r, err := l.client.do("GET", "/accounts/balance", "", true)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &balances); err != nil {
		return
	}
	return
}

// LongMarginMarket is
func (l *Liquid) LongMarginMarket(productID int, quantity float64) (response OrderResponse, err error) {
	// leverage_level is fixed at 2
	marginOrder := NewLeverageMarketOrder(productID, "buy", quantity, 2)
	type RequestOrder struct {
		Order LeverageMarketOrder `json:"order"`
	}
	requestOrder := RequestOrder{Order: marginOrder}
	jsonStr, err := json.Marshal(requestOrder)
	r, err := l.client.do("POST", "/orders/", string(jsonStr), true)
	if err != nil {
		return
	}
	fmt.Println("Created long market order")
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	return
}

// ShortMarginMarket is
func (l *Liquid) ShortMarginMarket(productID int, quantity float64) (response OrderResponse, err error) {
	// leverage_level is fixed at 2
	marginOrder := NewLeverageMarketOrder(productID, "sell", quantity, 2)
	type RequestOrder struct {
		Order LeverageMarketOrder `json:"order"`
	}
	requestOrder := RequestOrder{Order: marginOrder}
	jsonStr, err := json.Marshal(requestOrder)
	r, err := l.client.do("POST", "/orders/", string(jsonStr), true)
	if err != nil {
		return
	}
	fmt.Println("Created short market order")
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	return
}

func (l *Liquid) GetAnOrder(orderID int) (order OrderResponse, err error) {
	r, err := l.client.do("GET", "/orders/"+strconv.Itoa(orderID), "", true)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &order); err != nil {
		return
	}
	return
}

// GetTrades is
func (l *Liquid) GetTrades(fundingCurrency string, status string, tradingType string) (response TradeResponse, err error) {
	query := "/trades"
	params := []string{}
	if fundingCurrency != "" {
		params = append(params, "funding_currency="+fundingCurrency)
	}
	if status != "" {
		params = append(params, "status="+status)
	}
	if tradingType != "" {
		params = append(params, "trading_type="+tradingType)
	}
	if len(params) != 0 {
		query = query + "?" + strings.Join(params, "&")
	}
	r, err := l.client.do("GET", query, "", true)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	return
}

// UpdateTrade is
func (l *Liquid) UpdateTrade(tradeID int, stopLoss float64, takeProfit float64) (response OrderResponse, err error) {
	param := map[string]map[string]float64{
		"trade": map[string]float64{
			"stop_loss":   stopLoss,
			"take_profit": takeProfit,
		},
	}
	jsonStr, err := json.Marshal(param)
	r, err := l.client.do("PUT", fmt.Sprintf("/trades/%s", strconv.Itoa(tradeID)), string(jsonStr), true)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	return
}
