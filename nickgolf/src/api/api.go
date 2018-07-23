package api

import (
	"bmicalculate"
	"encoding/json"
	"net/http"
	"strconv"
)

type BmiStatus struct {
	Bmi    float64 `json:"bmi"`
	Status string  `json:"status"`
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	height := r.URL.Query().Get("height")
	weight := r.URL.Query().Get("weight")
	heightConv, _ := strconv.ParseFloat(height, 64)
	weightConv, _ := strconv.ParseFloat(weight, 64)

	bmiCalculate := bmicalculate.BmiCal(heightConv, weightConv)
	bmiStatusCheck := bmicalculate.BmiStatusCheck(bmiCalculate)

	bmiStatus := BmiStatus{Bmi: bmiCalculate, Status: bmiStatusCheck}
	bmiJson, _ := json.Marshal(bmiStatus)
	w.Write(bmiJson)
}
