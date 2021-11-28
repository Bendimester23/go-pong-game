package utils

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameObject interface {
	Update()
	Render()
	Reset()
}

func DrawTitleBtn(text string, id, selected int, offset float32) {
	alpha := uint8((Clamp(0, 30, offset+30.0-float32((id+6)*2)) / 30.0) * 256)
	if selected == id {
		DrawCenterText(text, int32(250+id*30+int(offset)), 20, rl.NewColor(0, 0, 0, alpha))
		rl.DrawLineEx(rl.NewVector2(float32(rl.GetScreenWidth()/2-int(rl.MeasureText(text, 20))/2), float32(270+id*30+int(offset))), rl.NewVector2(float32(rl.GetScreenWidth()/2+int(rl.MeasureText(text, 20))/2), float32(270+id*30+int(offset))), 2, rl.NewColor(0, 0, 0, alpha))
	} else {
		DrawCenterText(text, int32(250+id*30+int(offset)), 20, rl.NewColor(130, 130, 130, alpha))
	}
}

func DrawCenterText(text string, y int32, size int32, color rl.Color) {
	rl.DrawText(text, int32(rl.GetScreenWidth()/2-int(rl.MeasureText(text, size))/2), y, size, color)
}

func DrawCenterTextDef(text string, y int32, size int32) {
	rl.DrawText(text, int32(rl.GetScreenWidth()/2-int(rl.MeasureText(text, size))/2), y, size, rl.Gray)
}
