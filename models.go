package cloudpayments

import (
	"time"

	"github.com/guregu/null"
)

type Interval string

const (
	Day   Interval = "Day"
	Week  Interval = "Week"
	Month Interval = "Month"
)

type Currency string

const (
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type Model map[string]interface{}

type DefaultResponse struct {
	Success bool
	Message null.String
}

type PaymentResponse struct {
	DefaultResponse
	Model
}

type FindSubscriptionsResponse struct {
	DefaultResponse
	Model []Model
}

type CryptoPayRequest struct {
	Amount               float32
	Currency             Currency
	IpAddress            string
	CardCryptogramPacket string
	Name                 string
	PaymentUrl           string
	InvoiceId            string
	CultureName          string
	AccountId            string
	Email                string
}

type CreateSubscribeRequest struct {
	Token               string
	AccountId           string
	Description         string
	Email               string
	Amount              float32
	Currency            Currency
	RequireConfirmation bool
	StartDate           time.Time
	Interval            Interval
	Period              int
}

type SubscribeRequest struct {
	Id string
}

type FindSubscriptionsRequest struct {
	AccountId string
}

type UpdateSubscribeRequest struct {
	Id                  string
	Description         string
	Email               string
	Amount              float32
	Currency            Currency
	RequireConfirmation bool
	StartDate           time.Time
	Interval            Interval
	Period              int
}

type Post3dsRequest struct {
	TransactionId int
	PaRes         string
}
