package targetprocess

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestEncoders(t *testing.T) {
	Convey("Test Encoders", t, func() {
		Convey("BasicAuth to HASH", func() {
			user := "admin"
			pass := "pass"
			hash := "YWRtaW46cGFzcw=="
			testHash := encodeAuth(user, pass)
			So(testHash, ShouldEqual, hash)
		})
	})
}
