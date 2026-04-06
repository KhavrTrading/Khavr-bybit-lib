package bybit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
	"io"
	"net/http"
	"strconv"
)

const (
	timestampKey  = "X-BAPI-TIMESTAMP"
	signatureKey  = "X-BAPI-SIGN"
	apiRequestKey = "X-BAPI-API-KEY"
	recvWindowKey = "X-BAPI-RECV-WINDOW"
	signTypeKey   = "X-BAPI-SIGN-TYPE"
)

func (c *BybitClient) doGet(endpoint string, builder bybit_models.ParamBuilder, out interface{}) error {
	fullURL := c.BaseURL + endpoint
	queryString := BuildGetParams(builder.ToParams())
	fullURL += queryString

	timestamp := GetCurrentTime() // Subtract 2 seconds to account for network latency
	signature := c.SignPayload("GET", timestamp, "5000", nil, queryString)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create GET request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(signTypeKey, "2")
	req.Header.Set(apiRequestKey, c.APIKey)
	req.Header.Set(timestampKey, strconv.FormatInt(timestamp, 10))
	req.Header.Set(recvWindowKey, "5000")
	req.Header.Set(signatureKey, signature)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read GET response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", string(body))
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("failed to decode GET response: %w", err)
	}

	return nil
}

func (c *BybitClient) doPost(endpoint string, body interface{}, out interface{}) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal POST body: %w", err)
	}

	fullURL := c.BaseURL + endpoint
	timestamp := GetCurrentTime() // Subtract 2 seconds to account for network latency
	signature := c.SignPayload("POST", timestamp, "5000", bodyBytes, "")

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create POST request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(signTypeKey, "2")
	req.Header.Set(apiRequestKey, c.APIKey)
	req.Header.Set(timestampKey, strconv.FormatInt(timestamp, 10))
	req.Header.Set(recvWindowKey, "5000")
	req.Header.Set(signatureKey, signature)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read POST response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", string(respBytes))
	}

	if err := json.Unmarshal(respBytes, out); err != nil {
		return fmt.Errorf("failed to decode POST response: %w", err)
	}

	return nil
}
