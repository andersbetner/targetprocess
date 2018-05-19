package targetprocess

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestClientUtils(t *testing.T) {
	Convey("Client Util", t, func() {
		Convey("Generate TP Address", func() {
			Convey("with valid characters", func() {
				// passed values
				account := "testaccount"

				// expected values
				expectedAddress := "https://testaccount.tpondemand.com/api/v1/"

				// returned values
				returnedAddress, err := generateTPAddress(account)

				// run test
				So(err, ShouldBeNil)
				So(returnedAddress, ShouldEqual, expectedAddress)
			})
			Convey("With invalid characters", func() {
				// passed values
				account := "t estacco%unt"

				// returned values
				returnedAddress, err := generateTPAddress(account)

				// run test
				So(err, ShouldBeError)
				So(returnedAddress, ShouldBeEmpty)
			})
		})
	})
}
