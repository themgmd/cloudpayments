package cloudpayments

import (
	"encoding/json"
	"fmt"
)

func (cp *cloudPayments) PayByCryptogram(params CryptoPayRequest) (*PaymentResponse, error) {
	requestURL := cp.GenerateRequestURI(baseURL, paymentsGroup, cardsGroup, singlePayMethod)

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

func (cp *cloudPayments) Post3ds(params Post3dsRequest) (*PaymentResponse, error) {
	requestURL := cp.GenerateRequestURI(baseURL, paymentsGroup, cardsGroup, post3dsMethod)

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
