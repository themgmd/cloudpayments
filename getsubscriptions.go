package cloudpayments

import (
	"encoding/json"
	"fmt"
)

func (cp *cloudPayments) GetSubscription(params SubscribeRequest) (*PaymentResponse, error) {
	requestURL := cp.GenerateRequestURI(baseURL, subscriptionsGroup, getSubscriptionMethod)

	bt, err := json.Marshal(params)
	if err != nil {
		err = fmt.Errorf("json.Marshal: %w", err)
		return nil, err
	}

	resp, err := cp.GetPaymentResponse(requestURL, bt)
	if err != nil {
		return nil, fmt.Errorf("cp.GetPaymentResponse: %w", err)
	}

	return resp, nil
}
