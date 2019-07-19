package liquid

import "encoding/json"

type OrderBook struct {
	BuyPriceLevels  [][2]json.Number `json:"buy_price_levels"`
	SellPriceLevels [][2]json.Number `json:"sell_price_levels"`
}
