package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geraldofigueiredo/practice-go/testing/testable_code/services"
	"github.com/gin-gonic/gin"
)

type pingServiceMock struct {
	pingHandlerFn func() (string, error)
}

func (mock pingServiceMock) HandlePing() (string, error) {
	return mock.pingHandlerFn()
}

func TestPingWithError(t *testing.T) {
	serviceMock := pingServiceMock{}
	serviceMock.pingHandlerFn = func() (string, error) {
		return "", errors.New("error executing ping")
	}
	services.PingService = serviceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	Ping(context)

	if response.Code != http.StatusInternalServerError {
		t.Error("response code should be 500")
	}

	if response.Body.String() != "error executing ping" {
		t.Error("response body should say 'error'")
	}
}

func TestPingWithNoError(t *testing.T) {
	serviceMock := pingServiceMock{}
	serviceMock.pingHandlerFn = func() (string, error) {
		return "pong", nil
	}
	services.PingService = serviceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	Ping(context)

	if response.Code != http.StatusOK {
		t.Error("response code should be 200")
	}

	if response.Body.String() != "pong" {
		t.Error("response body should say 'pong'")
	}
}
