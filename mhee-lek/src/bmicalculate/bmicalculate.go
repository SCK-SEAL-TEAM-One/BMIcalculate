package bmicalculate

type BMIResponse struct {
	BMI    float64 `json:"bmi"`
	Status string  `json:"status"`
}

func BmiCalculate(height, weight int) BMIResponse {
	heightMetre := height / 100
	bmi := weight / (heightMetre * heightMetre)
	return BMIResponse{
		BMI: bmi,
	}
}
