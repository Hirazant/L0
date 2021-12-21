package model

type Order struct {
	OrderUid    string  `json:"order_uid"`
	TrackNumber string  `json:"track_number"`
	Payment     Payment `json:"payment"`
	Items       []Item  `json:"items"`
}

type Payment struct {
	Transaction string `json:"transaction"`
	RequestId   int    `json:"request_id"`
	Amount      int    `json:"amount"`
}

type Item struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
}
