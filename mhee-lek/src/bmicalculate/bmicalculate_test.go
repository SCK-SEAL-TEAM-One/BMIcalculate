package bmicalculate_test

import "testing"

func Test_BmiCalculate_Input_Height_170_And_Weight_80_Should_Be_Bmi_27dot7_And_Status_Fat(t *testing.T) {
	height := 170
	weight := 80
	expected := BMIResponse{
		BMI:    27.7,
		Status: "อ้วน",
	}

	actual := BmiCalculate(height, weight)

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func Test_getBmiStatus_Input_Bmi_27dot7_Should_Be_Fat(t *testing.T) {
	bmi := 27.7
	expected := "อ้วน"

	actual := getBmiStatus(bmi)

	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
