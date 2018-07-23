package bmi_test

import (
	. "bmi"
	"testing"
)

func Test_BmiCalculate_Input_Height_170_And_Weight_80_Should_Be_27_7(t *testing.T) {
	height := 170.0
	weight := 80.0
	expectedBmi := 27.7

	actualBmi := BmiCalculate(height, weight)

	if expectedBmi != actualBmi {
		t.Errorf("expectedbmi %f but got %f", expectedBmi, actualBmi)
	}
}

func Test_BmiStatus_Input_Bmi_27_7_Should_Be_Fat(t *testing.T) {
	bmi := 27.7
	expectedStatus := "อ้วน"
	actualStatus := BmiStatus(bmi)

	if expectedStatus != actualStatus {
		t.Errorf("expectedStatus %s but got %s", expectedStatus, actualStatus)
	}
}

func Test_Bmi_Input_Height_170_And_Weight_80_Should_Be_Bmi_27_7_And_Status_Fat(t *testing.T) {
	height := 170.0
	weight := 80.0
	expectedBmiResult := BmiResult{
		BMI:    27.7,
		Status: "อ้วน",
	}
	actualBmiResult := Bmi(height, weight)

	if expectedBmiResult != actualBmiResult {
		t.Errorf("expectedBmi %v but got %v", expectedBmiResult, actualBmiResult)
	}
}
