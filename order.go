package liquid

// NewMarketOrder is
func NewMarketOrder(productID int, side string, quantity float64) (order MarketOrder) {
	order = MarketOrder{
		ProductID: productID,
		OrderType: "market",
		Side:      side,
		Quantity:  quantity,
	}
	return
}

// NewLeverageMarketOrder is
func NewLeverageMarketOrder(productID int, side string, quantity float64, leverageLevel int) (order LeverageMarketOrder) {
	order = LeverageMarketOrder{
		MarketOrder:    NewMarketOrder(productID, side, quantity),
		LeverageLevel:  leverageLevel,
		OrderDirection: "two_direction",
	}
	return
}

// MarketOrder is ...
type MarketOrder struct {
	ProductID int     `json:"product_id"`
	OrderType string  `json:"order_type"`
	Side      string  `json:"side"`
	Quantity  float64 `json:"quantity"`
}

// LimitOrder is
type LimitOrder struct {
	ProductID int     `json:"product_id"`
	Side      string  `json:"side"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
}

// LeverageMarketOrder is
type LeverageMarketOrder struct {
	MarketOrder
	LeverageLevel  int    `json:"leverage_level"`
	OrderDirection string `json:"order_direction"`
}

// LeverageLimitOrder is
type LeverageLimitOrder struct {
	LimitOrder
	LeverageLevel  int    `json:"leverage_level"`
	OrderDirection string `json:"order_direction"`
}

// OrderResponse is
type OrderResponse struct {
	ID                   int         `json:"id"`
	OrderType            string      `json:"order_type"`
	Quantity             string      `json:"quantity"`
	DiscQuantity         string      `json:"disc_quantity"`
	IcebergTotalQuantity string      `json:"iceberg_total_quantity"`
	Side                 string      `json:"side"`
	FilledQuantity       string      `json:"filled_quantity"`
	Price                float64     `json:"price"`
	CreatedAt            int         `json:"created_at"`
	UpdatedAt            int         `json:"updated_at"`
	Status               string      `json:"status"`
	LeverageLevel        int         `json:"leverage_level"`
	SourceExchange       string      `json:"source_exchange"`
	ProductID            int         `json:"product_id"`
	ProductCode          string      `json:"product_code"`
	FundingCurrency      string      `json:"funding_currency"`
	CurrencyPairCode     string      `json:"currency_pair_code"`
	OrderFee             float64     `json:"order_fee"`
	Executions           []Execution `json:"executions"`
}

type Execution struct {
	ID        int    `json:"id"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
	TakerSide string `json:"taker_side"`
	MySide    string `json:"my_side"`
	CreatedAt int    `json:"created_at"`
}
