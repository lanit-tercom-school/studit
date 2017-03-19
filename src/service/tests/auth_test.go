package endpointTests

import (
	"testing"
	"service/auth"
	. "github.com/smartystreets/goconvey/convey"
)

type testPairInt struct {
	input int
	output int
}

func TestGenerateNewToken(t *testing.T) {
	t.Parallel()
	ns := []testPairInt{
		{-10, 0},
		{-1, 0},
		{0, 0},
		{1, 1},
		{10, 10},
	}
	for _, n := range ns {
		n := n
		Convey("Subject: generating random strings", t, func() {
			f := auth.GenerateNewToken(n.input)
			Convey("Length sould be equal", func() {
				So(len(f), ShouldEqual, n.output)
			})
		})
	}
}