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

	resp, err := cp.NewRequest(requestURL, bt)
	if err != nil {
		err = fmt.Errorf("cp.NewRequest: %w", err)
		return nil, err
	}

	var paymentResponse PaymentResponse
	if err = json.Unmarshal(resp, &paymentResponse); err != nil {
		err = fmt.Errorf("json.Unmarshal: %w", err)
		return nil, err
	}

	return &paymentResponse, nil
}
