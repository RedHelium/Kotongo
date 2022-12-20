package extensions

import (
	"strconv"
)

var (
	ClWhite = Color{R: 255, G: 255, B: 255, A: 255}
	ClBlack = Color{R: 0, G: 0, B: 0, A: 255}
)

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

func (color *Color) HEX() string {
	return Color2HEX(*color)
}

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
func RGB2HEX(red, green, blue float32) string {
	return Concat("#", FloatToHex(red), FloatToHex(green), FloatToHex(blue))
}

func Color2HEX(color Color) string {

	return Concat("#", FloatToHex(color.R), FloatToHex(color.G), FloatToHex(color.B))
}

func HEX2RGB(hex string) (red, green, blue float32) {

	r := ""
	g := ""
	b := ""
	rIndex := 1
	gIndex := 3
	bIndex := 5

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

	r1, _ := strconv.ParseInt(r, 16, 64)
	g1, _ := strconv.ParseInt(g, 16, 64)
	b1, _ := strconv.ParseInt(b, 16, 64)

	return float32(r1), float32(g1), float32(b1)
}

func HEX2Color(hex string) *Color {

	r, g, b := HEX2RGB(hex)

	return &Color{
		R: r,
		G: g,
		B: b,
		A: 255,
	}

}
