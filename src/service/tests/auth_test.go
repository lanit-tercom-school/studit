package endpointTests

import (
	"testing"
	"service/auth"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"net/http"
	"net/http/httptest"
	"github.com/astaxie/beego"
	"bytes"
)

type testPairInt struct {
	input int
	output int
}

func TestGenerateNewToken(t *testing.T) {
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


func TestLoginAndLogout(t *testing.T) {
	f, _ := json.Marshal(auth.Usr{Login:"a@a", Password:"a"})
	r, _ := http.NewRequest("POST", "http://localhost:8080/v1/auth/login", bytes.NewReader(f))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var response auth.LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Login\n", t, func() {
		Convey("Status code should be 200", func() {
			So(w.Code, ShouldEqual, 200)
			So(err, ShouldEqual, nil)
		})
	})

	r, _ = http.NewRequest("GET", "http://localhost:8080/v1/auth/logout/?token=" + response.Token, nil)
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Logout (normal)\n", t, func() {
		Convey("Status code should be 200", func() {
			So(w.Code, ShouldEqual, 200)
			So(w.Body.String(), ShouldEqual, "\"OK\"")
		})
	})
}

func TestEmptyLogout(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:8080/v1/auth/logout/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Logout (empty token)\n", t, func() {
		Convey("Status code should be 400", func() {
			So(w.Code, ShouldEqual, 400)
			So(w.Body.String(), ShouldEqual, "\"Empty token\"")
		})
	})
}

func getTestStrings() []string {
	var testStrings []string
	data, err := ioutil.ReadFile(`tests/blns.json`)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &testStrings)
	if err != nil {
		panic(err)
	}
	str := []string{"qwertyuiopasdfghjklzxcvbnm", "qwertyuiopasdfghjklzxcvbnm.s", "qwertyuiopasdfghjklzxcvbnm.",
		"qwertyuiopasdfghjklzxcvbnm.asdadsadsadsa.", "qwertyuiopasdfghjklzxcvbnm.sdfghjk.fghjklsa",
		"<>", "!", "@#", "$", "%", "^&*()`~", "#$%"}
	testStrings = append(testStrings, str...)
	return testStrings
}

func TestWrongTokenLogout(t *testing.T) {
	t.Parallel()
	testStrings := getTestStrings()
	for i, str := range testStrings {
		i := i
		str := str // capture range variable (from example on https://golang.org/pkg/testing/)

		Convey("String number " + strconv.Itoa(i) + " Sent wrong token " + str, t, func() {
			requestURL := "http://localhost:8080/v1/auth/logout/?token=" + str
			r, err := http.NewRequest("GET", requestURL, nil)
			if err != nil {
			} else {
				w := httptest.NewRecorder()
				beego.BeeApp.Handlers.ServeHTTP(w, r)
				Convey("Status code should be 400", func() {
					So(w.Code, ShouldEqual, 400)
				})
			}
		})
	}
}