package core

import "github.com/go-gl/mathgl/mgl64"

type Camera struct {
	position         mgl64.Vec3
	aspect           float32
	near             float32
	far              float32
	axis             int
	projection       int
	FOV              float32
	orthographicSize float32
	projectionMatrix mgl64.Mat4
}
