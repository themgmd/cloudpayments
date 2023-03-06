package cloudpayments

import (
	"bytes"
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
	client := &http.Client{}

	bodyReader := bytes.NewReader(body)
	
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		err = fmt.Errorf("http.NewRequest: %w", err)
		return nil, err
	}

	req.SetBasicAuth(cp.publicID, cp.apiSecret)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("client.Do: %w", err)
		return nil, err
	}

	bt, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("io.ReadAll: %w", err)
		return nil, err
	}

	return bt, nil
}
