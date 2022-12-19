package main

import (
	"kotongo/core"
	"kotongo/math"
)

func main() {

	backgroundColor := math.Color{R: 0, G: 0.5, B: 1.0, A: 1.0}
	core.InitWindow(1280, 720, "Test app", backgroundColor)
}
