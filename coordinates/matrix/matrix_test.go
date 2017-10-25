package matrix_test

import (
	"lengine/matrix"
	vector "lengine/vector"
	"math"
	"strconv"
	"testing"
)

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func TestOrthographic(t *testing.T) {
	v1 := vector.Vec4{512, 384, 1, 1}
	ortho := matrix.NewOrthographic(vector.Vec2{1024, 768})
	v1 = ortho.MultiplyVector(v1)

	if v1.X != 0 {
		errorString := "Expected X: 0 - Actual: " + strconv.FormatFloat(v1.X, 'f', -1, 64)
		t.Error(errorString)
	}
	if v1.Y != 0 {
		errorString := "Expected Y: 0 - Actual: " + strconv.FormatFloat(v1.Y, 'f', -1, 64)
		t.Error(errorString)
	}
	if v1.Z != -1 {
		errorString := "Expected Z: -1 - Actual: " + strconv.FormatFloat(v1.Z, 'f', -1, 64)
		t.Error(errorString)
	}
	if v1.W != 1 {
		errorString := "Expected W: 1 - Actual: " + strconv.FormatFloat(v1.W, 'f', -1, 64)
		t.Error(errorString)
	}
}

func TestTranslate(t *testing.T) {
	v1 := vector.Vec4{512, 384, 1, 1}
	mat := matrix.NewIdentity().Translate(vector.Vec3{1, 1, 0})
	v1 = mat.MultiplyVector(v1)

	if v1.X != 513 {
		errorString := "Expected X: 513 - Actual: " + strconv.FormatFloat(v1.X, 'f', -1, 64)
		t.Error(errorString)
	}
	if v1.Y != 385 {
		errorString := "Expected Y: 385 - Actual: " + strconv.FormatFloat(v1.Y, 'f', -1, 64)
		t.Error(errorString)
	}
}

func TestRotateZ(t *testing.T) {
	v1 := vector.Vec4{512, 384, 1, 1}
	mat := matrix.NewIdentity().RotateZ(180)
	v1 = mat.MultiplyVector(v1)

	if Round(v1.X) != -512 {
		errorString := "Expected X: -512 - Actual: " + strconv.FormatFloat(v1.X, 'f', -1, 64)
		t.Error(errorString)
	}
	if Round(v1.Y) != -384 {
		errorString := "Expected Y: -384 - Actual: " + strconv.FormatFloat(v1.Y, 'f', -1, 64)
		t.Error(errorString)
	}
}

func TestScaleAndInverse(t *testing.T) {
	v1 := vector.Vec4{2.5, 4, 1, 1}
	mat := matrix.NewIdentity().Scale(vector.Vec3{2, 2, 2})
	v1 = mat.MultiplyVector(v1)

	if Round(v1.X) != 5 {
		errorString := "Expected X: 5 - Actual: " + strconv.FormatFloat(v1.X, 'f', -1, 64)
		t.Error(errorString)
	}
	if Round(v1.Y) != 8 {
		errorString := "Expected Y: 8 - Actual: " + strconv.FormatFloat(v1.Y, 'f', -1, 64)
		t.Error(errorString)
	}
	if Round(v1.Z) != 2 {
		errorString := "Expected Z: 2 - Actual: " + strconv.FormatFloat(v1.Z, 'f', -1, 64)
		t.Error(errorString)
	}

	mat = mat.Inverse()
	v1 = mat.MultiplyVector(v1)

	if v1.X != 2.5 {
		errorString := "Expected X: 2.5 - Actual: " + strconv.FormatFloat(v1.X, 'f', -1, 64)
		t.Error(errorString)
	}
	if Round(v1.Y) != 4 {
		errorString := "Expected Y: 4 - Actual: " + strconv.FormatFloat(v1.Y, 'f', -1, 64)
		t.Error(errorString)
	}
	if Round(v1.Z) != 1 {
		errorString := "Expected Z: 1 - Actual: " + strconv.FormatFloat(v1.Z, 'f', -1, 64)
		t.Error(errorString)
	}
}
