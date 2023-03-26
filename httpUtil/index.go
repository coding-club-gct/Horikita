package httpUtil

import (
	"net"
	"net/http"
	"time"

	"github.com/joel-samuel-raj/Horikita/constants"
)

func CreateHTTPClientWithBearerToken() *http.Client {
	client := &http.Client{}
	client.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client.Transport = &headerTransport{
		transport: client.Transport,
		header: http.Header{
			"Authorization": []string{"Bearer " + constants.C.Strings["API_TOKEN"]},
		},
	}
	return client
}

type headerTransport struct {
	transport http.RoundTripper
	header    http.Header
}

func (ht *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header = ht.header.Clone()
	return ht.transport.RoundTrip(req)
}
