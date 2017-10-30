package matrix_test

import (
	"lengine/coordinates/matrix"
	vector "lengine/coordinates/vector"
	"lengine/testing/comparison"
	"testing"
)

func TestOrthographic(t *testing.T) {
	v1 := vector.Vec4{512, 384, 1, 1}
	ortho := matrix.NewOrthographic(vector.Vec2{1024, 768})
	v1 = ortho.MultiplyVector(v1)

	comparison.CompareEqualityFloat(0, v1.X, t)
	comparison.CompareEqualityFloat(0, v1.Y, t)
	comparison.CompareEqualityFloat(-1, v1.Z, t)
	comparison.CompareEqualityFloat(1, v1.W, t)
}

func TestTranslate(t *testing.T) {
	v1 := vector.Vec4{512, 384, 1, 1}
	mat := matrix.NewIdentity().Translate(vector.Vec3{1, 1, 0})
	v1 = mat.MultiplyVector(v1)

	comparison.CompareEqualityFloat(513, v1.X, t)
	comparison.CompareEqualityFloat(385, v1.Y, t)
}

func TestRotateZ(t *testing.T) {
	v1 := vector.Vec4{512, 384, 1, 1}
	mat := matrix.NewIdentity().RotateZ(180)
	v1 = mat.MultiplyVector(v1)

	comparison.CompareEqualityFloat(-512, comparison.Round(v1.X), t)
	comparison.CompareEqualityFloat(-384, comparison.Round(v1.Y), t)
}

func TestScaleAndInverse(t *testing.T) {
	v1 := vector.Vec4{2.5, 4, 1, 1}
	mat := matrix.NewIdentity().Scale(vector.Vec3{2, 2, 2})
	v1 = mat.MultiplyVector(v1)

	comparison.CompareEqualityFloat(5, comparison.Round(v1.X), t)
	comparison.CompareEqualityFloat(8, comparison.Round(v1.Y), t)
	comparison.CompareEqualityFloat(2, comparison.Round(v1.Z), t)

	mat = mat.Inverse()
	v1 = mat.MultiplyVector(v1)

	comparison.CompareEqualityFloat(2.5, v1.X, t)
	comparison.CompareEqualityFloat(4, v1.Y, t)
	comparison.CompareEqualityFloat(1, v1.Z, t)
}
