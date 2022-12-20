package pos

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Transform struct {
	position mgl64.Vec3
	rotation mgl64.Quat
	scale    mgl64.Vec3
}
