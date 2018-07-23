package api

import (
	"bmi"
	"encoding/json"
	"net/http"
	"strconv"
)

func BmiHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	heightText := query.Get("height")
	weightText := query.Get("weight")
	height, _ := strconv.ParseFloat(heightText, 64)
	weight, _ := strconv.ParseFloat(weightText, 64)
	result := bmi.Bmi(height, weight)
	resultJson, _ := json.Marshal(result)

	writer.Write(resultJson)
}
