package targetprocess

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestClientUtils(t *testing.T) {
	Convey("Test Utils", t, func() {
		Convey("Generate TP Address", func() {
			address := "https://md5.tpondemand.com/api/v1/"
			testAddress, err := generateTPAddress("md5")
			So(err, ShouldBeNil)
			So(testAddress, ShouldEqual, address)
		})
		Convey("Generate User Agent", func() {
			useragent := "targetprocess/development golang api client. running for account md5"
			testagent := generateUserAgent("md5")
			So(testagent, ShouldEqual, useragent)
		})
	})
}
