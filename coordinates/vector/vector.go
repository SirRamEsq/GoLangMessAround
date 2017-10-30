package vector

import "strconv"

//Vec2 represents a 2D coordinate pair
type Vec2 struct {
	X float64
	Y float64
}

//Vec3 represents a 3D coordinate pair
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (vecA Vec3) AddVec3(vecB Vec3) Vec3 {
	newVec := Vec3{}

	newVec.X = vecA.X + vecB.X
	newVec.Y = vecA.Y + vecB.Y
	newVec.Z = vecA.Z + vecB.Z
	newVec.X = vecA.X + vecB.X

	return newVec
}

func (vec Vec3) String() string {
	returnString := "X: '" + strconv.FormatFloat(vec.X, 'f', -1, 64) + "'\n"
	returnString += "Y: '" + strconv.FormatFloat(vec.Y, 'f', -1, 64) + "'\n"
	returnString += "Z: '" + strconv.FormatFloat(vec.Z, 'f', -1, 64) + "'\n"
	return returnString
}

//Vec4 represents a 4D coordinate pair
type Vec4 struct {
	X float64
	Y float64
	Z float64
	W float64
}
