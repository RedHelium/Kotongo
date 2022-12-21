package main

import (
	"kotongo/core"

	"github.com/RedHelium/toolz"
)

func main() {

	backgroundColor := toolz.HEX2Color("#1ec97c")

	core.InitWindow(1280, 720, "Test app", backgroundColor.Clamp01())
}
