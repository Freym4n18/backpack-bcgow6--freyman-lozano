package calculadora

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestRestar(t *testing.T) {
	a,b := 10, 5
	expected := 5

	result := restar(a,b)

	if result != expected {
		t.Errorf("expected: %d, got: %d", expected, result)
	}
}

func TestRestarTestify(t *testing.T) {
	a,b := 10, 5
	expected := 5

	result := restar(a,b)

	assert.Equal(t, result, expected, "expected: %d, got: %d", result, expected)
}

func TestSort(t *testing.T) {
	array := []int{3,2,4,7,1,6,5}
	expected := []int{1,2,3,4,5,6,7}

	result := Sort(array)

	assert.Equal(t, result, expected, "expected: %+v, got: %+v", result, expected)
}

func TestDividir(t *testing.T) {
	num,den := 12,3
	expected := 4
	result , err := Dividir(num,den)
	assert.Nil(t, err, "Error in Dividir")
	assert.Equal(t, result, expected, "expected: %d, got: %d", result, expected)
}