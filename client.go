package targetprocess

const (
	tpSchema  = "https://"
	tpBaseURL = ".tpondemand.com/api/v1/"
)

// NewBasicClient returns a new targetprocess api client
// using a username and password.
func NewBasicClient(account string, user string, pass string) (*TPClient, error) {
	client := new(TPClient)
	address, err := generateTPAddress(account)
	if err != nil {
		return client, err
	}
	client.url = address
	client.basicAuth = true
	client.auth = encodeAuth(user, pass)
	return client, nil
}

// NewTokenClient returns a new targetprocess api client
// using an auth token
func NewTokenClient(account string, token string) (*TPClient, error) {
	client := new(TPClient)
	address, err := generateTPAddress(account)
	if err != nil {
		return client, err
	}
	client.url = address
	client.tokenAuth = true
	client.auth = token
	return client, nil
}
