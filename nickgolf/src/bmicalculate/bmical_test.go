package bmicalculate_test

import (
	"bmicalculate"
	"testing"
)

func Test_BmiCal_Input_Height_170_And_Weight_80_Should_Be_27_Point_7(t *testing.T) {
	height := 170.00
	weight := 80.00
	expected := 27.7

	actual := bmicalculate.BmiCal(height, weight)

	if expected != actual {
		t.Errorf("expect %f but got %f", expected, actual)
	}

}

func Test_BmiStatusCheck_Input_27_Point_7_Should_Be_fat(t *testing.T) {
	bmiResult := 27.7
	expected := "อ้วน"

	actual := bmicalculate.BmiStatusCheck(bmiResult)

	if expected != actual {
		t.Errorf("expect %s but got %s", expected, actual)
	}
}
