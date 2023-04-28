package cloudpayments

const (
	HeaderXContentHMAC = "X-Content-HMAC"
	HeaderContentHMAC  = "Content-HMAC"

	baseURL = "api.cloudpayments.ru"

	paymentsGroup   = "payments"
	cardsGroup      = "cards"
	singlePayMethod = "charge"
	post3dsMethod   = "post3ds"

	subscriptionsGroup       = "subscriptions"
	createSubscriptionMethod = "create"
	getSubscriptionMethod    = "get"
	findSubscriptionMethod   = "find"
	updateSubscriptionMethod = "update"
	cancelSubscriptionMethod = "cancel"
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
