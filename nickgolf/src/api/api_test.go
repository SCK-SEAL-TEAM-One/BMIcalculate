package api_test

import (
	. "api"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ApiHandler_Input_Height_170_And_Weight_80_Should_Be_Json(t *testing.T) {

	request := httptest.NewRequest("GET", "/bmi?height=170&weight=80", nil)
	response := httptest.NewRecorder()

	expected := `{"bmi":27.7,"status":"อ้วน"}`

	ApiHandler(response, request)

	respBody := response.Result()
	actual, _ := ioutil.ReadAll(respBody.Body)

	if expected != string(actual) {
		t.Errorf("expect %s but got %s", expected, string(actual))
	}

}
