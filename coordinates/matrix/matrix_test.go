package matrix_test

import (
	"lengine/coordinates/matrix"
	vector "lengine/coordinates/vector"
	"math"
	"strconv"
	"testing"
)

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func CompareFloat(expected float64, actual float64, t *testing.T) {
	if expected != actual {
		errorString := "Expected X: " + strconv.FormatFloat(expected, 'f', -1, 64) + " - Actual: " + strconv.FormatFloat(actual, 'f', -1, 64)
		t.Error(errorString)
	}
}

func TestOrthographic(t *testing.T) {
	v1 := vector.Vec4{512, 384, 1, 1}
	ortho := matrix.NewOrthographic(vector.Vec2{1024, 768})
	v1 = ortho.MultiplyVector(v1)

	CompareFloat(0, v1.X, t)
	CompareFloat(0, v1.Y, t)
	CompareFloat(-1, v1.Z, t)
	CompareFloat(1, v1.W, t)
}

func TestTranslate(t *testing.T) {
	v1 := vector.Vec4{512, 384, 1, 1}
	mat := matrix.NewIdentity().Translate(vector.Vec3{1, 1, 0})
	v1 = mat.MultiplyVector(v1)

	CompareFloat(513, v1.X, t)
	CompareFloat(385, v1.Y, t)
}

func TestRotateZ(t *testing.T) {
	v1 := vector.Vec4{512, 384, 1, 1}
	mat := matrix.NewIdentity().RotateZ(180)
	v1 = mat.MultiplyVector(v1)

	CompareFloat(-512, Round(v1.X), t)
	CompareFloat(-384, Round(v1.Y), t)
}

func TestScaleAndInverse(t *testing.T) {
	v1 := vector.Vec4{2.5, 4, 1, 1}
	mat := matrix.NewIdentity().Scale(vector.Vec3{2, 2, 2})
	v1 = mat.MultiplyVector(v1)

	CompareFloat(5, Round(v1.X), t)
	CompareFloat(8, Round(v1.Y), t)
	CompareFloat(2, Round(v1.Z), t)

	mat = mat.Inverse()
	v1 = mat.MultiplyVector(v1)

	CompareFloat(2.5, v1.X, t)
	CompareFloat(4, v1.Y, t)
	CompareFloat(1, v1.Z, t)
}
