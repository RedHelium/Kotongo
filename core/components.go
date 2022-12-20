package core

import "kotongo/core/pos"

//Global object data
type Entity struct {
	name      string
	render    Renderer
	transform pos.Transform
}

//Render object data
type Renderer struct {
	drawable uint32
	points   []float32
}
