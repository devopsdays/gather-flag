package service

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/devopsdays/gather-flag/rooms/dbclient"
	"github.com/devopsdays/gather-flag/rooms/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRoom(t *testing.T) {
	mockRepo := &dbclient.MockMongoClient{}

	mockRepo.On("QueryRoom", "123").Return(model.Room{ID: "123", Roomname: "Room_123"}, nil)
	mockRepo.On("QueryRoom", "456").Return(model.Room{}, fmt.Errorf("Some error"))
	DBClient = mockRepo

	Convey("Given a HTTP request for /rooms/123", t, func() {
		req := httptest.NewRequest("GET", "/rooms/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)
				// So(resp.Body, ShouldEqual, "hi")

				// account := model.Room{}
				// json.Unmarshal(resp.Body.Bytes(), &account)
				// So(account, ShouldEqual, "hi")
				// So(account.ID, ShouldEqual, "123")
				// So(account.Roomname, ShouldEqual, "Person_123")
			})
		})
	})

	Convey("Given a HTTP request for /rooms/456", t, func() {
		req := httptest.NewRequest("GET", "/rooms/456", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestGetRoomWrongPath(t *testing.T) {

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
