package cloudpayments

import (
	"encoding/json"
	"fmt"
)

func (cp *cloudPayments) FindSubscriptions(params FindSubscriptionsRequest) (*FindSubscriptionsResponse, error) {
	requestURL := cp.GenerateRequestURI(BaseURL, SubscriptionsGroup, FindSubscriptionMethod)

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

	var findSubscriptionResponse FindSubscriptionsResponse
	if err = json.Unmarshal(resp, &findSubscriptionResponse); err != nil {
		err = fmt.Errorf("json.Unmarshal: %w", err)
		return nil, err
	}

	return &findSubscriptionResponse, nil
}
