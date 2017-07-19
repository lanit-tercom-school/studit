package endpointTests

import (
	"data-service/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewsModel(t *testing.T) {
	t_news := models.NewsJson{
		Title:       "TestNews1",
		Description: "Description of TestNews1",
		Tags:        []string{"Test", "Other"},
	}
	id, err := models.AddNews(&t_news)
	if err != nil {
		t.Error(err.Error())
	}
	t_news_t, err := models.GetNewsById(int(id))
	if err != nil {
		t.Error(err.Error())
	}
	Convey("Subject: Add; News should be the same", t, func() {
		So(t_news.Title, ShouldEqual, t_news_t.Title)
		So(t_news.Description, ShouldEqual, t_news_t.Description)
		So(len(t_news.Tags), ShouldEqual, len(t_news_t.Tags))
	})

	t_news_t.Title = "TestNews1_Changed"
	t_news_t.Description = "Changed description of TestNews1"
	err = models.UpdateNewsById(t_news_t)
	if err != nil {
		t.Error(err.Error())
	}

	t_news_t_t, err := models.GetNewsById(int(id))
	if err != nil {
		t.Error(err.Error())
	}
	Convey("Subject: Update; News should be the same", t, func() {
		So(t_news_t_t.Title, ShouldEqual, t_news_t.Title)
		So(t_news_t_t.Description, ShouldEqual, t_news_t.Description)
		So(len(t_news_t_t.Tags), ShouldEqual, len(t_news_t.Tags))
	})

	err = models.DeleteNews(t_news_t_t.Id)
	if err != nil {
		t.Error(err.Error())
	}
	t_news_t, err = models.GetNewsById(int(id))
	Convey("Subject: Delete; Error shouldn't be nil", t, func() {
		So(err, ShouldNotBeNil)
		So(t_news_t, ShouldBeNil)
	})
}

func TestNewsGetAllWithTag(t *testing.T) {
	tags := []string{"Other", "World", "School"}
	for _, tag := range tags {
		tag := tag // capture range variable (from example on https://golang.org/pkg/testing/)
		Convey("Subject: Get all news with filter by tag", t, func() {
			requestPath := "http://localhost:8080/v1/news/?tag=" + tag
			r, _ := http.NewRequest("GET", requestPath, nil)
			w := httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)

			var response []models.NewsJson
			err := json.Unmarshal(w.Body.Bytes(), &response)
			Convey("Response news should contain only news with current tag", func() {
				So(err, ShouldBeNil)
				for _, object := range response {
					So(models.TagInArrayOfStrings(tag, object.Tags), ShouldBeTrue)
				}
			})
		})

	}
}
