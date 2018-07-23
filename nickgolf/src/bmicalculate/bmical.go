package bmicalculate

import (
	"math"
)

func BmiCal(height, weight float64) float64 {
	bmiCalculate := weight / (height * height / 10000)
	return math.Ceil(bmiCalculate*10) / 10
}

func BmiStatusCheck(bmiResult float64) string {
	if bmiResult > 25.00 && bmiResult < 30.00 {
		return "อ้วน"
	}
	return ""
}
