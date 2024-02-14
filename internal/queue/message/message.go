package message

type Message interface {
	GetRoutingKey() string
}

const EnrollRoutingKey = "enrollment"

type EnrollMessage struct {
	RoutingKey    string `json:"-"`
	TransactionId int64  `json:"transaction_id"`
}

func (m *EnrollMessage) GetRoutingKey() string {
	return m.RoutingKey
}
