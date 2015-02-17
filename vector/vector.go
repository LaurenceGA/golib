/*
Package vector defines methods and structs for manipulating vectors
*/
package vector

import "math"

//Vector2 is a 2 value vector
type Vector2 [2]float64

//Vector3 is a 3 value vector
type Vector3 [3]float64

//Vector4 is a 4 value vector
type Vector4 [4]float64

/////////////////Addition/////////////////

//Add combines two Vector2 types
func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{v1[0] + v2[0], v1[1] + v2[1]}
}

//Add combines two Vector3 types
func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

//Add combines two Vector4 types
func (v1 Vector4) Add(v2 Vector4) Vector4 {
	return Vector4{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

/////////////////Subtraction/////////////////

//Sub subtracts two Vector2 types
func (v1 Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{v1[0] - v2[0], v1[1] - v2[1]}
}

//Sub combines two Vector3 types
func (v1 Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

//Sub combines two Vector4 types
func (v1 Vector4) Sub(v2 Vector4) Vector4 {
	return Vector4{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

/////////////////Multiplication/////////////////

//Mul multiplies the vector by a constant
func (v1 Vector2) Mul(c float64) Vector2 {
	return Vector2{v1[0] * c, v1[1] * c}
}

//Mul multiplies the vector by a constant
func (v1 Vector3) Mul(c float64) Vector3 {
	return Vector3{v1[0] * c, v1[1] * c, v1[2] * c}
}

//Mul multiplies the vector by a constant
func (v1 Vector4) Mul(c float64) Vector4 {
	return Vector4{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

/////////////////Len/////////////////

//Len returns the length of the vector
//This is the square root of the sum of the squares
func (v1 Vector2) Len() float64 {
	return float64(math.Hypot(v1[0], v1[1]))
}

//Len returns the length of the vector
//This is the square root of the sum of the squares
func (v1 Vector3) Len() float64 {
	return float64(math.Sqrt((v1[0] * v1[0]) + (v1[1] * v1[1]) + (v1[2] * v1[2])))
}

//Len returns the length of the vector
//This is the square root of the sum of the squares
func (v1 Vector4) Len() float64 {
	return float64(math.Sqrt((v1[0] * v1[0]) + (v1[1] * v1[1]) + (v1[2] * v1[2]) + (v1[3] * v1[3])))
}

/////////////////Normalize/////////////////

//Normalize sets the vector magnitude to 1
func (v1 Vector2) Normalize() Vector2 {
	l := 1.0 / v1.Len()
	return Vector2{v1[0] * l, v1[1] * l}
}

//Normalize sets the vector magnitude to 1
func (v1 Vector3) Normalize() Vector3 {
	l := 1.0 / v1.Len()
	return Vector3{v1[0] * l, v1[1] * l, v1[2] * l}
}

//Normalize sets the vector magnitude to 1
func (v1 Vector4) Normalize() Vector4 {
	l := 1.0 / v1.Len()
	return Vector4{v1[0] * l, v1[1] * l, v1[2] * l, v1[3] * l}
}

/////////////////XYZW Gets/////////////////

////Vector2////

//X will return vector element mapped to xyzw
func (v1 Vector2) X() float64 {
	return v1[0]
}

//Y  will return vector element mapped to xyzw
func (v1 Vector2) Y() float64 {
	return v1[1]
}

////Vector3////

//X will return vector element mapped to xyzw
func (v1 Vector3) X() float64 {
	return v1[0]
}

//Y will return vector element mapped to xyzw
func (v1 Vector3) Y() float64 {
	return v1[1]
}

//Z will return vector element mapped to xyzw
func (v1 Vector3) Z() float64 {
	return v1[2]
}

////Vector4////

//X will return vector element mapped to xyzw
func (v1 Vector4) X() float64 {
	return v1[0]
}

//Y will return vector element mapped to xyzw
func (v1 Vector4) Y() float64 {
	return v1[1]
}

//Z will return vector element mapped to xyzw
func (v1 Vector4) Z() float64 {
	return v1[2]
}

//W will return vector element mapped to xyzw
func (v1 Vector4) W() float64 {
	return v1[3]
}
