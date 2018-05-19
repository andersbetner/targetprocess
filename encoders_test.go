package targetprocess

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestEncoders(t *testing.T) {
	Convey("Encoding a given value", t, func() {
		Convey("of a username and password to Base64", func() {
			// passed values
			username := "test"
			password := "test"

			// expected values
			expectedB64String := "dGVzdDp0ZXN0"

			// run test
			b64String := encodeAuth(username, password)

			So(b64String, ShouldEqual, expectedB64String)
		})
	})
}
