package x_youtube

import (
	"errors"
	"net/http"
)

type TokenKey struct {
	ApiKey      string
	AccessToken string

	// Transport is the underlying HTTP transport.
	// If nil, http.DefaultTransport is used.
	Transport http.RoundTripper
}

func (t *TokenKey) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
		if rt == nil {
			return nil, errors.New("googleapi/transport: no Transport specified or available")
		}
	}
	newReq := *req
	newReq.Header.Add("Authorization", "Bearer "+t.AccessToken)
	args := newReq.URL.Query()
	args.Set("key", t.ApiKey)
	newReq.URL.RawQuery = args.Encode()

	return rt.RoundTrip(&newReq)
}
