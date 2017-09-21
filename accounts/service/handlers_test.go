package service

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/devopsdays/gather-flag/accounts/dbclient"
	"github.com/devopsdays/gather-flag/accounts/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAccount(t *testing.T) {
	mockRepo := &dbclient.MockMongoClient{}

	mockRepo.On("QueryAccount", "123").Return(model.Account{ID: "123", Username: "Person_123"}, nil)
	mockRepo.On("QueryAccount", "456").Return(model.Account{}, fmt.Errorf("Some error"))
	DBClient = mockRepo

	Convey("Given a HTTP request for /accounts/123", t, func() {
		req := httptest.NewRequest("GET", "/accounts/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)
				// So(resp.Body, ShouldEqual, "hi")

				// account := model.Account{}
				// json.Unmarshal(resp.Body.Bytes(), &account)
				// So(account, ShouldEqual, "hi")
				// So(account.ID, ShouldEqual, "123")
				// So(account.Username, ShouldEqual, "Person_123")
			})
		})
	})

	Convey("Given a HTTP request for /accounts/456", t, func() {
		req := httptest.NewRequest("GET", "/accounts/456", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestGetAccountWrongPath(t *testing.T) {

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
