package httpUtil

import (
	"net"
	"net/http"
	"time"

	"github.com/joel-samuel-raj/Horikita/constants"
)

func CreateHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
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
		},
	}
}

func AddAuthorizationHeader(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+constants.C.Strings["API_TOKEN"])
}
