package matrix

import (
	vector "lengine/vector"
	"math"
)

const (
	pi          = 3.14159265
	oneDegInRad = (2.0 * pi) / 360.0 // 0.017444444
	oneRadInDeg = 360.0 / (2.0 * pi) //57.2957795
)

/*Matrix represents a 4x4 matrix
 * stored like this:
 * 	0  4  8   12
 * 	1  5  9   13
 * 	2  6  10  14
 * 	3  7  11  15
 */
type Matrix struct {
	values [16]float64
}

//New Retruns a new Matrix
func New() Matrix {
	return Matrix{}
}

//NewOrthographic returns a new Matrix with an orthographic perspective
func NewOrthographic(viewSize vector.Vec2) Matrix {
	var orthX float64
	var orthY float64
	orthX = 2.0 / viewSize.X
	orthY = 2.0 / -viewSize.Y

	return Matrix{
		[16]float64{orthX, 0.0, 0.0, 0.0,
			0.0, orthY, 0.0, 0.0,
			0.0, 0.0, -1.0, 0.0,
			-1.0, 1.0, 0.0, 1.0}}
}

//NewIdentity returns a new Identity Matrix
func NewIdentity() Matrix {
	return Matrix{
		[16]float64{1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1}}
}

//MultiplyVector multiplies a Vec4 and a matrix and returns the resulting Vec4
func (m Matrix) MultiplyVector(v vector.Vec4) vector.Vec4 {
	// 0x + 4y + 8z + 12w
	x :=
		m.values[0]*v.X +
			m.values[4]*v.Y +
			m.values[8]*v.Z +
			m.values[12]*v.W
	// 1x + 5y + 9z + 13w
	y :=
		m.values[1]*v.X +
			m.values[5]*v.Y +
			m.values[9]*v.Z +
			m.values[13]*v.W
	// 2x + 6y + 10z + 14w
	z :=
		m.values[2]*v.X +
			m.values[6]*v.Y +
			m.values[10]*v.Z +
			m.values[14]*v.W
	// 3x + 7y + 11z + 15w
	w :=
		m.values[3]*v.X +
			m.values[7]*v.Y +
			m.values[11]*v.Z +
			m.values[15]*v.W
	return vector.Vec4{X: x, Y: y, Z: z, W: w}
}

//MultiplyMatrix multiplies two matricies together
func (m Matrix) MultiplyMatrix(mm Matrix) Matrix {
	newMat := Matrix{}
	rIndex := 0
	var sum float64
	for col := 0; col < 4; col++ {
		for row := 0; row < 4; row++ {
			sum = 0.0
			for i := 0; i < 4; i++ {
				sum += mm.values[i+col*4] * m.values[row+i*4]
			}
			newMat.values[rIndex] = sum
			rIndex++
		}
	}
	return newMat
}

//Translate translates the matrix by a Vec3
func (m Matrix) Translate(v vector.Vec3) Matrix {
	mTranslate := NewIdentity()
	mTranslate.values[12] = v.X
	mTranslate.values[13] = v.Y
	mTranslate.values[14] = v.Z
	return mTranslate.MultiplyMatrix(m)
}

//RotateZ rotates the matrix by a Vec3
func (m Matrix) RotateZ(deg float64) Matrix {
	// convert to radians
	rad := deg * oneDegInRad
	mRotate := NewIdentity()
	mRotate.values[0] = math.Cos(rad)
	mRotate.values[4] = -math.Sin(rad)
	mRotate.values[1] = math.Sin(rad)
	mRotate.values[5] = math.Cos(rad)
	return mRotate.MultiplyMatrix(m)
}

//Scale scales the matrix by a Vec3
func (m Matrix) Scale(scale vector.Vec3) Matrix {
	mScale := NewIdentity()
	mScale.values[0] = scale.X
	mScale.values[5] = scale.Y
	mScale.values[10] = scale.Z
	return mScale.MultiplyMatrix(m)
}

//Determinant returns the determinant of a matrix
func (m Matrix) Determinant() float64 {
	return m.values[12]*m.values[9]*m.values[6]*m.values[3] -
		m.values[8]*m.values[13]*m.values[6]*m.values[3] -
		m.values[12]*m.values[5]*m.values[10]*m.values[3] +
		m.values[4]*m.values[13]*m.values[10]*m.values[3] +
		m.values[8]*m.values[5]*m.values[14]*m.values[3] -
		m.values[4]*m.values[9]*m.values[14]*m.values[3] -
		m.values[12]*m.values[9]*m.values[2]*m.values[7] +
		m.values[8]*m.values[13]*m.values[2]*m.values[7] +
		m.values[12]*m.values[1]*m.values[10]*m.values[7] -
		m.values[0]*m.values[13]*m.values[10]*m.values[7] -
		m.values[8]*m.values[1]*m.values[14]*m.values[7] +
		m.values[0]*m.values[9]*m.values[14]*m.values[7] +
		m.values[12]*m.values[5]*m.values[2]*m.values[11] -
		m.values[4]*m.values[13]*m.values[2]*m.values[11] -
		m.values[12]*m.values[1]*m.values[6]*m.values[11] +
		m.values[0]*m.values[13]*m.values[6]*m.values[11] +
		m.values[4]*m.values[1]*m.values[14]*m.values[11] -
		m.values[0]*m.values[5]*m.values[14]*m.values[11] -
		m.values[8]*m.values[5]*m.values[2]*m.values[15] +
		m.values[4]*m.values[9]*m.values[2]*m.values[15] +
		m.values[8]*m.values[1]*m.values[6]*m.values[15] -
		m.values[0]*m.values[9]*m.values[6]*m.values[15] -
		m.values[4]*m.values[1]*m.values[10]*m.values[15] +
		m.values[0]*m.values[5]*m.values[10]*m.values[15]
}

//Transpose does something
func (m Matrix) Transpose() Matrix {
	return Matrix{
		[16]float64{m.values[0], m.values[4], m.values[8], m.values[12],
			m.values[1], m.values[5], m.values[9], m.values[13],
			m.values[2], m.values[6], m.values[10], m.values[14],
			m.values[3], m.values[7], m.values[11], m.values[15]}}
}

//Inverse returns the inverse of a matrix
func (m Matrix) Inverse() Matrix {
	det := m.Determinant()
	/* there is no inverse if determinant is zero (not likely unless scale is
	broken) */
	if 0.0 == det {
		//Matrix has no determinant. can not invert
		return NewIdentity()
	}
	invDet := 1.0 / det

	return Matrix{
		[16]float64{
			invDet * (m.values[9]*m.values[14]*m.values[7] - m.values[13]*m.values[10]*m.values[7] +
				m.values[13]*m.values[6]*m.values[11] - m.values[5]*m.values[14]*m.values[11] -
				m.values[9]*m.values[6]*m.values[15] + m.values[5]*m.values[10]*m.values[15]),
			invDet * (m.values[13]*m.values[10]*m.values[3] - m.values[9]*m.values[14]*m.values[3] -
				m.values[13]*m.values[2]*m.values[11] + m.values[1]*m.values[14]*m.values[11] +
				m.values[9]*m.values[2]*m.values[15] - m.values[1]*m.values[10]*m.values[15]),
			invDet * (m.values[5]*m.values[14]*m.values[3] - m.values[13]*m.values[6]*m.values[3] +
				m.values[13]*m.values[2]*m.values[7] - m.values[1]*m.values[14]*m.values[7] -
				m.values[5]*m.values[2]*m.values[15] + m.values[1]*m.values[6]*m.values[15]),
			invDet * (m.values[9]*m.values[6]*m.values[3] - m.values[5]*m.values[10]*m.values[3] -
				m.values[9]*m.values[2]*m.values[7] + m.values[1]*m.values[10]*m.values[7] +
				m.values[5]*m.values[2]*m.values[11] - m.values[1]*m.values[6]*m.values[11]),
			invDet * (m.values[12]*m.values[10]*m.values[7] - m.values[8]*m.values[14]*m.values[7] -
				m.values[12]*m.values[6]*m.values[11] + m.values[4]*m.values[14]*m.values[11] +
				m.values[8]*m.values[6]*m.values[15] - m.values[4]*m.values[10]*m.values[15]),
			invDet * (m.values[8]*m.values[14]*m.values[3] - m.values[12]*m.values[10]*m.values[3] +
				m.values[12]*m.values[2]*m.values[11] - m.values[0]*m.values[14]*m.values[11] -
				m.values[8]*m.values[2]*m.values[15] + m.values[0]*m.values[10]*m.values[15]),
			invDet * (m.values[12]*m.values[6]*m.values[3] - m.values[4]*m.values[14]*m.values[3] -
				m.values[12]*m.values[2]*m.values[7] + m.values[0]*m.values[14]*m.values[7] +
				m.values[4]*m.values[2]*m.values[15] - m.values[0]*m.values[6]*m.values[15]),
			invDet * (m.values[4]*m.values[10]*m.values[3] - m.values[8]*m.values[6]*m.values[3] +
				m.values[8]*m.values[2]*m.values[7] - m.values[0]*m.values[10]*m.values[7] -
				m.values[4]*m.values[2]*m.values[11] + m.values[0]*m.values[6]*m.values[11]),
			invDet * (m.values[8]*m.values[13]*m.values[7] - m.values[12]*m.values[9]*m.values[7] +
				m.values[12]*m.values[5]*m.values[11] - m.values[4]*m.values[13]*m.values[11] -
				m.values[8]*m.values[5]*m.values[15] + m.values[4]*m.values[9]*m.values[15]),
			invDet * (m.values[12]*m.values[9]*m.values[3] - m.values[8]*m.values[13]*m.values[3] -
				m.values[12]*m.values[1]*m.values[11] + m.values[0]*m.values[13]*m.values[11] +
				m.values[8]*m.values[1]*m.values[15] - m.values[0]*m.values[9]*m.values[15]),
			invDet * (m.values[4]*m.values[13]*m.values[3] - m.values[12]*m.values[5]*m.values[3] +
				m.values[12]*m.values[1]*m.values[7] - m.values[0]*m.values[13]*m.values[7] -
				m.values[4]*m.values[1]*m.values[15] + m.values[0]*m.values[5]*m.values[15]),
			invDet * (m.values[8]*m.values[5]*m.values[3] - m.values[4]*m.values[9]*m.values[3] -
				m.values[8]*m.values[1]*m.values[7] + m.values[0]*m.values[9]*m.values[7] +
				m.values[4]*m.values[1]*m.values[11] - m.values[0]*m.values[5]*m.values[11]),
			invDet * (m.values[12]*m.values[9]*m.values[6] - m.values[8]*m.values[13]*m.values[6] -
				m.values[12]*m.values[5]*m.values[10] + m.values[4]*m.values[13]*m.values[10] +
				m.values[8]*m.values[5]*m.values[14] - m.values[4]*m.values[9]*m.values[14]),
			invDet * (m.values[8]*m.values[13]*m.values[2] - m.values[12]*m.values[9]*m.values[2] +
				m.values[12]*m.values[1]*m.values[10] - m.values[0]*m.values[13]*m.values[10] -
				m.values[8]*m.values[1]*m.values[14] + m.values[0]*m.values[9]*m.values[14]),
			invDet * (m.values[12]*m.values[5]*m.values[2] - m.values[4]*m.values[13]*m.values[2] -
				m.values[12]*m.values[1]*m.values[6] + m.values[0]*m.values[13]*m.values[6] +
				m.values[4]*m.values[1]*m.values[14] - m.values[0]*m.values[5]*m.values[14]),
			invDet * (m.values[4]*m.values[9]*m.values[2] - m.values[8]*m.values[5]*m.values[2] +
				m.values[8]*m.values[1]*m.values[6] - m.values[0]*m.values[9]*m.values[6] -
				m.values[4]*m.values[1]*m.values[10] + m.values[0]*m.values[5]*m.values[10])}}
}
