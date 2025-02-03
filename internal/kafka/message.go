package kafka

import "fmt"

const OrderStatusTopic string = "order_status"

type OrderStatus struct {
	OrderId int    `json:"order_id"`
	Status  string `json:"status"`
}

func (o OrderStatus) ToJson() []byte {
	return []byte(fmt.Sprintf(`{"order_id":%d,"status":"%s"}`, o.OrderId, o.Status))
}
