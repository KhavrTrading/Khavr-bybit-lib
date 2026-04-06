package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/time/rate"
)

type BybitClient struct {
	APIKey    string
	APISecret string

	HTTPClient *http.Client
	BaseURL    string

	ipLimiter *rate.Limiter

	mu       sync.Mutex
	limiters map[string]*rate.Limiter
}

type countingTransport struct {
	base           *http.Transport
	reqCount       uint32
	resetThreshold uint32
}

func newCountingTransport(threshold uint32) *countingTransport {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
		MaxIdleConns:    100,
		IdleConnTimeout: 30 * time.Second,
	}
	http2.ConfigureTransport(tr)
	return &countingTransport{
		base:           tr,
		resetThreshold: threshold,
	}
}

func (c *countingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := c.base.RoundTrip(req)
	// increment *after* the request so we get a full stream count
	if cnt := atomic.AddUint32(&c.reqCount, 1); cnt >= c.resetThreshold {
		// gracefully close the HTTP/2 session
		c.base.CloseIdleConnections()
		// reset the counter so we do this periodically
		atomic.StoreUint32(&c.reqCount, 0)
	}
	return resp, err
}

func NewBybitClient(apiKey, apiSecret string) *BybitClient {

	ct := newCountingTransport(15_000)

	return &BybitClient{
		APIKey:    apiKey,
		APISecret: apiSecret,
		HTTPClient: &http.Client{
			Transport: ct,
			Timeout:   10 * time.Second,
		},
		BaseURL:   "https://api.bybit.com",
		ipLimiter: rate.NewLimiter(rate.Limit(20), 20),
		limiters:  make(map[string]*rate.Limiter),
	}
}

func (c *BybitClient) SignPayload(method string, timeStamp int64, recvWindow string, body []byte, queryString string) string {
	var suffix string
	if method == "POST" {
		suffix = string(body)
	} else {
		suffix = queryString[1:]
	}

	payload := strconv.FormatInt(timeStamp, 10) + c.APIKey + recvWindow + suffix

	h := hmac.New(sha256.New, []byte(c.APISecret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *BybitClient) GetLimiter(uid string) *rate.Limiter {
	c.mu.Lock()
	defer c.mu.Unlock()

	limiter, ok := c.limiters[uid]
	if !ok {
		limiter = rate.NewLimiter(rate.Limit(10), 10)
		c.limiters[uid] = limiter
	}
	return limiter
}
