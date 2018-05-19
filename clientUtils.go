package targetprocess

import "net/url"

func generateTPAddress(account string) (string, error) {
	address := tpSchema + account + tpBaseURL
	_, err := url.ParseRequestURI(address)
	if err != nil {
		return "", err
	}
	return address, nil
}
