package comparison

import (
	"math"
	"strconv"
	"testing"
)

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func CompareEqualityFloat(expected float64, actual float64, t *testing.T) {
	if expected != actual {
		errorString := "Expected: " + strconv.FormatFloat(expected, 'f', -1, 64) + " - Actual: " + strconv.FormatFloat(actual, 'f', -1, 64)
		t.Error(errorString)
	}
}

func CompareInequalityFloat(expected float64, actual float64, t *testing.T) {
	if expected == actual {
		errorString := "Expected NOT: " + strconv.FormatFloat(expected, 'f', -1, 64) + " - Actual: " + strconv.FormatFloat(actual, 'f', -1, 64)
		t.Error(errorString)
	}
}

func CompareEqualityBool(expected bool, actual bool, t *testing.T) {
	if expected != actual {
		errorString := "Expected: " + strconv.FormatBool(expected) + " - Actual: " + strconv.FormatBool(actual)
		t.Error(errorString)
	}
}

func CompareInequalityBool(expected bool, actual bool, t *testing.T) {
	if expected == actual {
		errorString := "Expected NOT: " + strconv.FormatBool(expected) + " - Actual: " + strconv.FormatBool(actual)
		t.Error(errorString)
	}
}
