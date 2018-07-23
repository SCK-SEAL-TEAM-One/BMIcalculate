package api

import (
	"encoding/json"
	"net/http"
)

type BMIResponse struct {
	BMI    float64 `json:"bmi"`
	Status string  `json:"status"`
}

func BMIHandler(write http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	height, _ := queryString.Get("height")
	weight, _ := queryString.Get("weight")
	bmi := BmiCalculate(height, weight)
	bmiJson, _ := json.Marshal(bmi)
	write.Write(bmiJson)
}
