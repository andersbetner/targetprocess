package targetprocess

import (
	"fmt"
	"net/url"
)

func generateTPAddress(account string) (string, error) {
	address := tpSchema + account + tpBaseURL
	_, err := url.ParseRequestURI(address)
	if err != nil {
		return "", err
	}
	return address, nil
}

func generateUserAgent(account string) string {
	return fmt.Sprintf("targetprocess/%s golang api client. running for account %s", version, account)
}
