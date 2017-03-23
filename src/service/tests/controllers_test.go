package endpointTests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"service/models"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"runtime"
	"path/filepath"
	_ "service/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"service/auth"
	"bytes"
	"github.com/vetcher/jwt"
	"time"
	"strconv"
	"io/ioutil"
)

func init() {
	jwt.GlobalStorage = jwt.NewStorage(time.Hour)
	_, file, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(path)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
}

// Tests /landing_page METHODS
// should return 3 Project structures
func TestLandingPageGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:8080/v1/land/projects", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var response []models.Project

	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Landing page GET method\n", t, func() {
		Convey("Status code should be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The result should not be empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		/*Convey("It should contain 3 items", func() {
			So(len(response), ShouldEqual, 0)
		})*/
		Convey("There should be a result for db_init_data.go", func() {
			So(response[0].Id, ShouldEqual, 1)
			So(response[0].Name, ShouldEqual, "Образовательный портал Studit")
			So(response[0].Logo, ShouldEqual, "/logo/1.jpg")
			So(response[0].Description, ShouldEqual, "Разработка образовательного портала для Lanit-Tercom School")

			So(response[1].Id, ShouldEqual, 2)
			So(response[1].Name, ShouldEqual, "Модный фрилансер")
			So(response[1].Logo, ShouldEqual, "/logo/2.jpg")
			So(response[1].Description, ShouldEqual, "Какие же стрелочки вокруг ноубука!")

			So(response[2].Id, ShouldEqual, 3)
			So(response[2].Name, ShouldEqual, "Оригинальное название")
			So(response[2].Logo, ShouldEqual, "/logo/3.jpg")
			So(response[2].Description, ShouldEqual, "Click-bait описание")
		})
	})
}

func TestLandingPagePut(t *testing.T) {
	r, _ := http.NewRequest("PUT", "http://localhost:8080/v1/land/projects/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Landing page PUT method\n" + r.URL.String(), t, func() {
		Convey("Status code should be 404", func() {
			So(w.Code, ShouldEqual, 404)
			So(w.Body.String(), ShouldEqual, "\"Not Found\"")
		})
	})
}

func TestLandingPagePost(t *testing.T) {
	r, _ := http.NewRequest("POST", "http://localhost:8080/v1/land/projects/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Landing page POST method\n", t, func() {
		Convey("Status code should be 404", func() {
			So(w.Code, ShouldEqual, 404)
			So(w.Body.String(), ShouldEqual, "\"Not Found\"")
		})
	})
}

func TestLandingPageDelete(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "http://localhost:8080/v1/land/projects/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Landing page DELETE method\n", t, func() {
		Convey("Status code should be 404", func() {
			So(w.Code, ShouldEqual, 404)
			So(w.Body.String(), ShouldEqual, "\"Not Found\"")
		})
	})
}

func TestLoginAndLogout(t *testing.T) {
	f, _ := json.Marshal(auth.Usr{Login:"a@a", Password:"a"})
	r, _ := http.NewRequest("POST", "http://localhost:8080/v1/auth/login", bytes.NewReader(f))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var response auth.UserAndToken
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