package cloudpayments

import (
	"encoding/json"
	"fmt"
)

func (cp *cloudPayments) UpdateSubscribe(params UpdateSubscribeRequest) (*PaymentResponse, error) {
	requestURL := cp.GenerateRequestURI(baseURL, subscriptionsGroup, updateSubscriptionMethod)

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
