package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type Vector2f struct {
	X float32
	Y float32
}

type Vector2i struct {
	X int32
	Y int32
}

//Source: https://easings.net/#easeInOutQuad
func EaseInOutQuad(x float32) float32 {
	if x < 0.5 {
		return 2 * x * x
	} else {
		return 1 - (-2*x+2)*(-2*x+2)/2
	}
}

func Clamp(min, max, v float32) float32 {
	if v < min {
		return min
	} else if v > max {
		return max
	}
	return v
}

func GrayWithAlpha(a, offset float32) rl.Color {

	return rl.NewColor(100, 100, 100, GetAlpha(a, offset))
}

func GetAlpha(a, offset float32) uint8 {
	if a > -1 {
		return 255
	}
	return uint8((Clamp(0, 30, a+30.0-offset) / 30.0) * 256)
}
