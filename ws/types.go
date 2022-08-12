package ws

type Arg struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

type ReceiveData struct {
	Arg  Arg           `json:"arg"`
	Data []interface{} `json:"data"`
}

type TickerData struct {
	LastPrice    string `json:"last"`
	LastSize     string `json:"lastSz"`
	BestAskPrice string `json:"askPx,"`
	BestAskSize  string `json:"askSz"`
	BestBidPrice string `json:"bidPx"`
	BestBidSize  string `json:"bidSz"`
	TimeStamp    int64  `json:"ts,string"`
}

type Ticker struct {
	Arg  Arg          `json:"arg"`
	Data []TickerData `json:"data"`
}

type TradesData struct {
	InstId    string `json:"instId"`
	TradeId   int64  `json:"tradeId,string"`
	Price     string `json:"px"`
	Size      string `json:"sz"`
	Side      string `json:"side"`
	TimeStamp int64  `json:"ts,string"`
}

type Trades struct {
	Arg  Arg          `json:"arg"`
	Data []TradesData `json:"data"`
}

type OrderbookData struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	TimeStamp int64      `json:"ts,string"`
	CheckSum  int64      `json:"checksum"`
}

type Orderbook struct {
	Arg    Arg             `json:"arg"`
	Action string          `json:"action"`
	Data   []OrderbookData `json:"data"`
}
