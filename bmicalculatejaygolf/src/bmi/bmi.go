package bmi

import "math"

type BmiResult struct {
	BMI    float64 `json:"bmi"`
	Status string  `json:"status"`
}

func Bmi(height, weight float64) BmiResult {
	bmi := BmiCalculate(height, weight)
	status := BmiStatus(bmi)

	return BmiResult{
		BMI:    bmi,
		Status: status,
	}
}
func BmiCalculate(height, weight float64) float64 {
	heightmeter := height / 100
	bmi := weight / (heightmeter * heightmeter)
	return math.Round(bmi*10) / 10
}

func BmiStatus(bmi float64) string {
	if bmi <= 18.5 {
		return "ผอม"
	}
	if bmi >= 18.6 && bmi <= 22.9 {
		return "น้ำหนักปกติ"
	}
	if bmi >= 23.0 && bmi <= 24.9 {
		return "น้ำหนักเกิน"
	}
	if bmi >= 25.0 && bmi <= 29.9 {
		return "อ้วน"
	}
	return "อ้วนมาก"
}
