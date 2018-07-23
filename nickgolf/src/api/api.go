package api

import (
	"encoding/json"
	"net/http"
)

type BmiStatus struct {
	Bmi    float64 `json:"bmi"`
	Status string  `json:"status"`
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	bmiStatus := BmiStatus{Bmi: 27.7, Status: "อ้วน"}
	bmiJson, _ := json.Marshal(bmiStatus)
	w.Write(bmiJson)
}
