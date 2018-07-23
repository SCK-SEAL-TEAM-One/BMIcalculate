package api_test

import (
	. "api"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_BMIHandler_Input_height_170_And_weight_80_Should_Be_BMI_27_7_And_Status_Fat(t *testing.T) {

	request := httptest.NewRequest("GET", "/bmi/", nil)
	responseRecorder := httptest.NewRecorder()
	expected := BMIResponse{
		BMI:    27.7,
		Status: "อ้วน",
	}
	BMIHandler(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	var actual BMIResponse
	json.Unmarshal(body, &actual)

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
