package service

import (
	"encoding/json"
	"fmt"
	"github.com/TonyXMH/goblog/accountservice/dbclient"
	"github.com/TonyXMH/goblog/accountservice/model"
	"github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
)

func TestGetAccount(t *testing.T) {
	mockRepo:=&dbclient.MockBoltClient{}
	mockRepo.On("QueryAccount","123").Return(model.Account{ID:"123",Name:"Person_123",},nil)
	mockRepo.On("QueryAccount","456").Return(model.Account{},fmt.Errorf("some error"))
	DBClient = mockRepo
	convey.Convey("Given a HTTP request for /accounts/123",t, func() {
		req:=httptest.NewRequest("GET","/accounts/123",nil)
		resp:=httptest.NewRecorder()
		convey.Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp,req)
			convey.Convey("Then the response should be 200", func() {
				convey.So(resp.Code,convey.ShouldEqual,200)
				account:=model.Account{}
				json.Unmarshal(resp.Body.Bytes(),&account)
				convey.So(account.ID,convey.ShouldEqual,"123")
				convey.So(account.Name,convey.ShouldEqual,"Person_123")
			})
		})
	})

	convey.Convey("Given a HTTP request for /accounts/456",t, func() {
		req:=httptest.NewRequest("GET","/accounts/456",nil)
		resp:=httptest.NewRecorder()
		convey.Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp,req)
			convey.Convey("Then the response should be a 404",func(){
				convey.So(resp.Code,convey.ShouldEqual,404)
			})
		})
	})
}

func TestGetAccountWrongPath(t *testing.T) {
	convey.Convey("Given a HTTP request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()
		convey.Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)
			convey.Convey("Then the response should be a 404", func() {
				convey.So(resp.Code, convey.ShouldEqual, 404)
			})
		})
	})

}
