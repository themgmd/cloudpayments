package cloudpayments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (cp *cloudPayments) GenerateRequestURI(baseURL string, parts ...string) string {
	var sb strings.Builder
	sb.WriteString(baseURL + "/")
	for index, part := range parts {
		sb.WriteString(part)
		if index != len(parts)-1 {
			sb.WriteString("/")
		}
	}

	return sb.String()
}

func (cp *cloudPayments) NewRequest(url string, body []byte) ([]byte, error) {
	var (
		client     = &http.Client{}
		bodyReader = bytes.NewReader(body)
	)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		err = fmt.Errorf("http.NewRequest: %w", err)
		return nil, err
	}

	req.SetBasicAuth(cp.publicID, cp.apiSecret)

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do: %w", err)
		return nil, err
	}

	defer resp.Body.Close()

	bt, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("io.ReadAll: %w", err)
		return nil, err
	}

	return bt, nil
}

func (cp *cloudPayments) GetPaymentResponse(requestURL string, data []byte) (*PaymentResponse, error) {
	resp, err := cp.NewRequest(requestURL, data)
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
