package model

type Order struct {
	OrderUid    string  `json:"order_uid"`
	TrackNumber string  `json:"track_number"`
	Payment     Payment `json:"payment"`
	Items       []Item  `json:"items"`
}

//type Cache interface {
//	Get(key, val interface{}) error
//}

//type OrderService struct {
//cache Cache
//}

//func (o *OrderService) Print() error {
//	var order Order
//	if err := o.cache.Get(1, &order); err != nil {
//		return err
//	}
//	fmt.Print(order)
//	return nil
//}

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

type EventStore struct {
	AggregateId   string `protobuf:"bytes,1,opt,name=aggregate_id,json=aggregateId" json:"aggregate_id,omitempty"`
	AggregateType string `protobuf:"bytes,2,opt,name=aggregate_type,json=aggregateType" json:"aggregate_type,omitempty"`
	EventId       string `protobuf:"bytes,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	EventType     string `protobuf:"bytes,4,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	EventData     string `protobuf:"bytes,5,opt,name=event_data,json=eventData" json:"event_data,omitempty"`
}
