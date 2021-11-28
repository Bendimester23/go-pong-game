package actors

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bar struct {
	X, Y    int32
	Height  int32
	Width   int32
	KeyUp   int32
	KeyDown int32
	Speed   int32
	Color   rl.Color
}

func (b *Bar) Reset() {
	b.Y = 0
}

func (b *Bar) Update() {
	if rl.IsKeyDown(b.KeyUp) && b.Y > 0 {
		b.Y -= b.Speed
	} else if rl.IsKeyDown(b.KeyDown) && b.Y < 400 {
		b.Y += b.Speed
	}
}

func (b *Bar) Render() {
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, b.Color)
}
