package cloudpayments

const (
	BaseURL = "api.cloudpayments.ru"

	PaymentsGroup   = "payments"
	CardsGroup      = "cards"
	SinglePayMethod = "charge"

	SubscriptionsGroup       = "subscriptions"
	CreateSubscriptionMethod = "create"
	GetSubscriptionMethod    = "get"
	FindSubscriptionMethod   = "find"
	UpdateSubscriptionMethod = "update"
	CancelSubscriptionMethod = "cancel"
)

type CloudPayments interface {
	SetSecret(secret string)
	SetPublicID(pubID string)

	PayByCryptogram(params CryptoPayRequest) (*PaymentResponse, error)
	Subscribe(params CreateSubscribeRequest) (*PaymentResponse, error)
	UpdateSubscribe(params UpdateSubscribeRequest) (*PaymentResponse, error)
	GetSubscription(params SubscribeRequest) (*PaymentResponse, error)
	FindSubscriptions(params FindSubscriptionsRequest) (*FindSubscriptionsResponse, error)
	Unsubscribe(params SubscribeRequest) (*PaymentResponse, error)
}

type cloudPayments struct {
	apiSecret string
	publicID  string
}

func New(secret, publicID string) CloudPayments {
	return &cloudPayments{secret, publicID}
}

func (cp *cloudPayments) SetSecret(secret string) {
	cp.apiSecret = secret
}

func (cp *cloudPayments) SetPublicID(pubID string) {
	cp.publicID = pubID
}
