package extensions

import (
	"strconv"
)

// TODO Add some colors
var (
	ClWhite = Color{R: 255, G: 255, B: 255, A: 255}
	ClBlack = Color{R: 0, G: 0, B: 0, A: 255}
)

// Basic color struct
type Color struct {
	R float32
	G float32
	B float32
	A float32
}

// Get color HEX code
func (color *Color) HEX() string {
	return Color2HEX(*color)
}

// Get color R, G, B, A values in range [0;1]
func (color *Color) Clamp01() *Color {
	return &Color{
		R: color.R / 255,
		G: color.G / 255,
		B: color.B / 255,
		A: color.A / 255,
	}
}

// TODO Replace in github project
// TODO Add support alpha channel

// Get HEX code from red, green and blue color channels
func RGB2HEX(red, green, blue float32) string {
	return Concat("#", FloatToHex(red), FloatToHex(green), FloatToHex(blue))
}

// Convert color struct to HEX code
func Color2HEX(color Color) string {

	return Concat("#", FloatToHex(color.R), FloatToHex(color.G), FloatToHex(color.B))
}

// TODO Test with zero values
// Get red, green and blue color values from HEX code
func HEX2RGB(hex string) (red, green, blue float32) {

	r := ""
	g := ""
	b := ""
	rIndex := 1
	gIndex := 3
	bIndex := 5

	//Check HEX code on zero values
	//////////////////
	if string(hex[rIndex]) != "0" {
		r = string(hex[rIndex : rIndex+2])
	} else {
		r = "0"
		gIndex--
		bIndex--
	}

	if string(hex[gIndex]) != "0" {
		g = string(hex[gIndex : gIndex+2])
	} else {
		g = "0"
		bIndex--
	}

	if string(hex[bIndex]) != "0" {
		b = string(hex[bIndex : bIndex+2])
	} else {
		b = "0"
	}
	//////////////////

	r1, _ := strconv.ParseInt(r, 16, 64)
	g1, _ := strconv.ParseInt(g, 16, 64)
	b1, _ := strconv.ParseInt(b, 16, 64)

	return float32(r1), float32(g1), float32(b1)
}

// Get color struct from HEX code
func HEX2Color(hex string) *Color {

	r, g, b := HEX2RGB(hex)

	return &Color{
		R: r,
		G: g,
		B: b,
		A: 255,
	}

}
