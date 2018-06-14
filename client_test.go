package targetprocess

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClients(t *testing.T) {
	Convey("Creating a new client", t, func() {
		account := "md5"
		username := "test"
		password := "test"
		authToken := "somerandomtoken"
		apiToken := "someapitoken"
		Convey("Basic Auth Client", func() {
			client, err := NewBasicClient(account, username, password)
			So(err, ShouldBeNil)
			So(client, ShouldNotBeNil)
		})
		Convey("Token Auth Client", func() {
			client, err := NewAuthTokenClient(account, authToken)
			So(err, ShouldBeNil)
			So(client, ShouldNotBeNil)
		})
		Convey("API Token Client", func() {
			client, err := NewTokenClient(account, apiToken)
			So(err, ShouldBeNil)
			So(client, ShouldNotBeNil)
		})
	})
}

func TestClientPrivate(t *testing.T) {
	Convey("Test Client util functionality", t, func() {
		account := "md5"
		token := "someTokenHere"
		md5HostAddress := "https://md5.tpondemand.com/api/v1/"
		headers := make(http.Header)
		headers["Test"] = []string{"header"}
		Convey("getHTTPClient check", func() {
			httpClient := getHTTPClient()
			So(httpClient, ShouldNotBeNil)
			So(httpClient, ShouldHaveSameTypeAs, &http.Client{})
		})
		Convey("newClient check", func() {
			testClient, err := newClient(account, token)
			So(err, ShouldBeNil)
			So(testClient, ShouldNotBeNil)
			So(testClient, ShouldNotBeEmpty)
			So(testClient.Token, ShouldEqual, token)
			So(testClient.HostAddress, ShouldEqual, md5HostAddress)
			So(testClient.Header, ShouldHaveSameTypeAs, http.Header{})
			So(testClient.scheme, ShouldEqual, "https://")
			So(testClient.httpClient, ShouldHaveSameTypeAs, &http.Client{})
			So(testClient.httpClient, ShouldNotBeNil)
			So(testClient.httpClient, ShouldNotBeEmpty)
			So(testClient.userAgent, ShouldNotBeNil)
			So(testClient.userAgent, ShouldNotBeEmpty)
		})
		Convey("SetHeader check", func() {
			client, _ := newClient(account, token)
			client.SetHeader("test", "header")
			So(client.Header, ShouldResemble, headers)
		})
	})
}
