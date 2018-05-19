package targetprocess

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestClients(t *testing.T) {
	Convey("Creating a client of the type", t, func() {
		Convey("BasicAuth", func() {
			// passed values
			account := "restapi"
			badAccount := "restapi "
			username := "John"
			password := "123"

			// expected values
			testB64String := encodeAuth(username, password)
			testURL, _ := generateTPAddress(account)
			basicTPClient := TPClient{
				url:       testURL,
				auth:      testB64String,
				basicAuth: true,
			}
			badBasicTPClient := TPClient{}
			// run tests now
			Convey("with valid args", func() {
				basicAuthClient, err := NewBasicClient(account, username, password)
				So(err, ShouldBeNil)
				So(basicAuthClient, ShouldResemble, &basicTPClient)
			})
			Convey("with invalid args", func() {
				basicAuthClient, err := NewBasicClient(badAccount, username, password)
				So(err, ShouldBeError)
				So(basicAuthClient, ShouldResemble, &badBasicTPClient)
			})
		})
		Convey("TokenAuth", func() {
			// passed values
			account := "restapi"
			badAccount := "rest api"
			token := "somestrangetoken"

			// expected values
			testURL, _ := generateTPAddress(account)
			tokenTPClient := TPClient{
				url:       testURL,
				auth:      token,
				tokenAuth: true,
			}
			badTokenTPClient := TPClient{}
			// run tests
			Convey("with valid args", func() {
				tokenAuthClient, err := NewTokenClient(account, token)
				So(err, ShouldBeNil)
				So(tokenAuthClient, ShouldResemble, &tokenTPClient)
			})
			Convey("with invalid args", func() {
				Convey("with valid args", func() {
					tokenAuthClient, err := NewTokenClient(badAccount, token)
					So(err, ShouldBeError)
					So(tokenAuthClient, ShouldResemble, &badTokenTPClient)
				})
			})
		})
	})
}
