package main

import (
	"kotongo/core"
	"kotongo/extensions"
)

func main() {

	backgroundColor := extensions.HEX2Color("#1ec97c")

	core.InitWindow(1280, 720, "Test app", backgroundColor.Clamp01())
}
