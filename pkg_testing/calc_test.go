package calc_test

import (
	"github.com/AndreoBouzas/estudos_golang/calculator/calc"
	"testing"
)

type TestCase struct {
	n1       int
	n2       int
	sig      string
	expected int
	recived  int
}

func main() {
	testeee := calc.calculator(10, 20, "-")
	fmt.Println(testeee)
}
func TestCalculator(t *testing.T) {
	testCase := TestCase{
		n1:       10,
		n2:       10,
		sig:      "+",
		expected: 20,
	}

	testCase.recived = calc.calculator(testCase.n1, testCase.n2, testCase.sig)
	if testCase.recived != testCase.expected {
		t.Fail()
	}
}
