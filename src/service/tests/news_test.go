package endpointTests

import (
	"testing"
	"service/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewsModel(t *testing.T) {
	t_news := models.NewsJson{
		Title: "TestNews1",
		Description: "Description of TestNews1",
		Tags: []string{"Test", "Other"},
	}
	id, err := models.AddNews(&t_news)
	if err != nil {
		t.Error(err.Error())
	}
	t_news_t, err := models.GetNewsById(int(id))
	if err != nil {
		t.Error(err.Error())
	}
	Convey("Subject: Add; News should be the same", t, func(){
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
	Convey("Subject: Update; News should be the same", t, func(){
		So(t_news_t_t.Title, ShouldEqual, t_news_t.Title)
		So(t_news_t_t.Description, ShouldEqual, t_news_t.Description)
		So(len(t_news_t_t.Tags), ShouldEqual, len(t_news_t.Tags))
	})

	err = models.DeleteNews(t_news_t_t.Id)
	if err != nil {
		t.Error(err.Error())
	}
	t_news_t, err = models.GetNewsById(int(id))
	Convey("Subject: Delete; Error shouldn't be nil", t, func(){
		So(err, ShouldNotBeNil)
		So(t_news_t, ShouldBeNil)
	})
}
