package shapes

import (
	"fmt"
	"kotongo/math"
	"reflect"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shape struct {
	drawable uint32
	points   []float32
	position mgl32.Vec3
	scale    mgl32.Vec3
}

var (
	TrianglePoints = []math.Vector3{
		{X: 0, Y: 0.1, Z: 0},     // top
		{X: -0.1, Y: -0.1, Z: 0}, // left
		{X: 0.1, Y: -0.1, Z: 0},  // right
	}
)

func Create(points []math.Vector3, pos, size mgl32.Vec3) *Shape {

	allPoints := make([]float32, len(points))

	for _, vect := range points {
		for i := 0; i < reflect.ValueOf(vect).NumField(); i++ {
			allPoints = append(allPoints, float32(reflect.ValueOf(vect).Field(i).Float()))
		}
	}
	//TODO Make a set position and size
	for i := range allPoints {
		/* if i == 0 || i == 3 || i == 6 {
			allPoints[i] += pos.X()
			allPoints[i] *= size.X()
		} else if i == 1 || i == 4 || i == 7 {
			allPoints[i] += pos.Y() * size.Y()
			allPoints[i] *= size.Y()
		} else if i == 2 || i == 5 || i == 8 {
			allPoints[i] += pos.Z() * size.Z()
			allPoints[i] *= size.Z()
		}
		*/
		fmt.Println("Index =", i, "Point = ", allPoints[i])
	}

	makePoints := make([]float32, len(allPoints))
	copy(makePoints, allPoints)

	return &Shape{
		drawable: MakeVAO(makePoints),
		points:   makePoints,
		position: pos,
		scale:    size,
	}
}

// Create vertex array object
func MakeVAO(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

// Draw shape vertex array
func (shape *Shape) Draw() {

	gl.BindVertexArray(shape.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(shape.points)/3))
}
