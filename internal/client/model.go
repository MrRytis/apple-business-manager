package client

type BulkEnrollRequest struct {
	RequestContext RequestContext `json:"requestContext"`
	TransactionId  string         `json:"transactionId"`
	DepResellerId  string         `json:"depResellerId"`
	Orders         []*Order       `json:"orders"`
}

type RequestContext struct {
	ShipTo   string `json:"shipTo"`
	LangCode string `json:"langCode"`
	TimeZone string `json:"timeZone"`
}

type Order struct {
	OrderNumber string `json:"orderNumber"`
	OrderDate   string `json:"orderDate"`
	OrderType   string `json:"orderType"`
	CustomerId  string `json:"customerId"`
	Deliveries  []*Delivery
}

type Delivery struct {
	DeliveryNumber string `json:"deliveryNumber"`
	ShipDate       string `json:"shipDate"`
	Devices        []*Device
}

type Device struct {
	Imei string `json:"deviceId"`
}

type BulkEnrollResponse struct {
	AppleTransactionId string             `json:"deviceEnrollmentTransactionId"`
	Status             BulkStatusResponse `json:"enrollDevicesResponse"`
}

type BulkStatusResponse struct {
	Code string `json:"statusCode"`
	Msg  string `json:"statusMessage"`
}

type BulkEnrollErrorResponse struct {
	ErrCode       string `json:"errorCode"`
	ErrMsg        string `json:"errorMessage"`
	TransactionId string `json:"transactionId"`
}

type StatusRequest struct {
	RequestContext     RequestContext `json:"requestContext"`
	DepResellerId      string         `json:"depResellerId"`
	AppleTransactionId string         `json:"deviceEnrollmentTransactionId"`
}

type StatusResponse struct {
}
