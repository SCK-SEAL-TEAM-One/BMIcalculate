package api_test

import (
    "testing"
    "net/http/httptest"
    "io/ioutil"
    "encoding/json"
)

func Test_BMIHandler_Input_height_170_And_weight_80_Should_Be_BMI_27_7_And_Status_Fat(t *testing.T) {

    request := httptest.NewRequest("GET", "/bmi/", nil)
    responseRecorder := httptest.NewRecorder()
    expected := BMIResponse{
		BMI: 27.7,
		Status: "อ้วน",
	}
    BMIHandler(request,responseRecorder)

    response := responseRecorder.Result()
    body,_ := ioutil.ReadAll(response.Body)

    var actual BMIResponse
    json.Unmarshal(body,&actual)

    if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}