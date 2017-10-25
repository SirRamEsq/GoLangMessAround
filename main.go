package main

import (
	"fmt"
	matrix "lengine/matrix"
	vector "lengine/vector"
)

func main() {
	v1 := vector.Vec4{X: 512, Y: 384, Z: 1, W: 1}
	ortho := matrix.NewOrthographic(vector.Vec2{X: 1024, Y: 768})
	v1 = ortho.MultiplyVector(v1)

	fmt.Println("TEST")
	fmt.Println(v1.X)
	fmt.Println(v1.Y)
	fmt.Println(v1.Z)
	fmt.Println(v1.W)
}
