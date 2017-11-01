package comparison

import (
	"lengine/entity"
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

func CompareEqualityString(expected string, actual string, t *testing.T) {
	if expected != actual {
		errorString := "Expected: " + expected + " - Actual: " + actual
		t.Error(errorString)
	}
}

func CompareInequalityString(expected string, actual string, t *testing.T) {
	if expected == actual {
		errorString := "Expected NOT: " + expected + " - Actual: " + actual
		t.Error(errorString)
	}
}

func CompareEqualityEID(expected entity.EID, actual entity.EID, t *testing.T) {
	if expected != actual {
		errorString := "Expected: " + expected.String() + " - Actual: " + actual.String()
		t.Error(errorString)
	}
}

func CompareInequalityEID(expected entity.EID, actual entity.EID, t *testing.T) {
	if expected == actual {
		errorString := "Expected NOT: " + expected.String() + " - Actual: " + actual.String()
		t.Error(errorString)
	}
}
