package service

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/devopsdays/gather-flag/topics/dbclient"
	"github.com/devopsdays/gather-flag/topics/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTopic(t *testing.T) {
	mockRepo := &dbclient.MockMongoClient{}

	mockRepo.On("QueryTopic", "123").Return(model.Topic{ID: "123", Title: "Title_123"}, nil)
	mockRepo.On("QueryTopic", "456").Return(model.Topic{}, fmt.Errorf("Some error"))
	DBClient = mockRepo

	Convey("Given a HTTP request for /topics/123", t, func() {
		req := httptest.NewRequest("GET", "/topics/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)
				// So(resp.Body, ShouldEqual, "hi")

				// topic := model.Topic{}
				// json.Unmarshal(resp.Body.Bytes(), &topic)
				// So(topic, ShouldEqual, "hi")
				// So(topic.ID, ShouldEqual, "123")
				// So(topic.Username, ShouldEqual, "Person_123")
			})
		})
	})

	Convey("Given a HTTP request for /topics/456", t, func() {
		req := httptest.NewRequest("GET", "/topics/456", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestGetTopicWrongPath(t *testing.T) {

	Convey("Given a HTTP request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}
