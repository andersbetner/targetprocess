package targetprocess

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

const (
	version   = "development"
	tpSchema  = "https://"
	tpBaseURL = ".tpondemand.com/api/v1/"
)

var (
	// SSLTimeout specifies how long a HTTP client can wait for a SSL handshake to
	// to complete
	SSLTimeout = 10
	// ConnectionTimeout specifies how a http client will wait for a response from
	// the target server
	ConnectionTimeout = 20
	// HTTPKeepAlive specified how long a http client will keep a connection open
	// with the target server. This can help reduce the time spent makeing calls
	// to the target server in the case of subsecuent calls
	HTTPKeepAlive = 30
)

// TPClient contains the base client for accessing the API
type TPClient struct {
	HostAddress string
	Token       string
	Header      http.Header

	httpClient  *http.Client
	scheme      string
	userAgent   string
	isBasicAuth bool
	isAuthToken bool
	isToken     bool
}

// NewBasicClient returns a new targetprocess api client
// using a username and password.
func NewBasicClient(account string, user string, pass string) (*TPClient, error) {
	token := encodeAuth(user, pass)
	client, err := newClient(account, token)
	if err != nil {
		return nil, err
	}
	client.isBasicAuth = true
	client.SetHeader("Authorization", (fmt.Sprintf("Basic %s", client.Token)))
	return client, nil
}

// NewAuthTokenClient return a new client using an Auth Token
func NewAuthTokenClient(account, token string) (*TPClient, error) {
	client, err := newClient(account, token)
	if err != nil {
		return nil, err
	}
	client.isAuthToken = true
	return client, err
}

// NewTokenClient returns a new client using a api token
func NewTokenClient(account, token string) (*TPClient, error) {
	client, err := newClient(account, token)
	if err != nil {
		return nil, err
	}
	client.isToken = true
	return client, nil
}

// SetHeader allows headers to manually be set
func (c *TPClient) SetHeader(header string, value string) {
	c.Header.Add(header, value)
}

// func (c *TPClient) AddQueryParams() {}

func newClient(account string, token string) (*TPClient, error) {
	HostAddress, err := generateTPAddress(account)
	if err != nil {
		return nil, err
	}
	client := &TPClient{
		HostAddress: HostAddress,
		Token:       token,
		Header:      make(http.Header),
		httpClient:  getHTTPClient(),
		scheme:      tpSchema,
		userAgent:   generateUserAgent(account),
	}
	return client, nil
}

func getHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   time.Duration(ConnectionTimeout) * time.Second,
				KeepAlive: time.Duration(HTTPKeepAlive) * time.Second,
			}).Dial,
			TLSHandshakeTimeout: time.Duration(SSLTimeout) * time.Second,
		},
	}
	return client
}
