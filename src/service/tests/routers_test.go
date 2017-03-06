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
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
}

type ErrorResponseType struct {
	Error string `json:"error"`
}

// Tests /landing_page METHODS
// should return 3 Project structures
func TestLandingPageGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:8080/v1/land_projects", nil)
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
	r, _ := http.NewRequest("PUT", "http://localhost:8080/v1/land_projects/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var response ErrorResponseType

	json.Unmarshal(w.Body.Bytes(), &response)
	Convey("Subject: Landing page PUT method\n" + r.URL.String(), t, func() {
		Convey("Status code should be 405", func() {
			So(w.Code, ShouldEqual, 405)
			So(response.Error, ShouldEqual, "Method Not Allowed")
		})
	})
}

func TestLandingPagePost(t *testing.T) {
	r, _ := http.NewRequest("POST", "http://localhost:8080/v1/land_projects/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var response ErrorResponseType

	json.Unmarshal(w.Body.Bytes(), &response)
	Convey("Subject: Landing page POST method\n", t, func() {
		Convey("Status code should be 405", func() {
			So(w.Code, ShouldEqual, 405)
			So(response.Error, ShouldEqual, "Method Not Allowed")
		})
	})
}

func TestLandingPageDelete(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "http://localhost:8080/v1/land_projects/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var response ErrorResponseType

	json.Unmarshal(w.Body.Bytes(), &response)
	Convey("Subject: Landing page DELETE method\n", t, func() {
		Convey("Status code should be 405", func() {
			So(w.Code, ShouldEqual, 405)
			So(response.Error, ShouldEqual, "Method Not Allowed")
		})
	})
}

