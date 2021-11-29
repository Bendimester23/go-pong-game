package actors

import (
	"math/rand"
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Position rl.Vector2
	Velocity utils.Vector2f
	Sat      float32
	LeftBar  *Bar
	RightBar *Bar
	Speed    float32
}

func (b *Ball) Reset() {
	b.Position = rl.NewVector2(float32(rand.Intn(600)+100), float32(rand.Intn(300)+50))
	b.Velocity = utils.Vector2f{X: b.Speed, Y: b.Speed}
	b.Sat = 0
}

func (b *Ball) Update() {
	b.Position.X += b.Velocity.X
	b.Position.Y += b.Velocity.Y

	if utils.Save.WeirdMode {
		if b.Velocity.Y < 0 {
			b.Velocity.X *= 1.01
		} else {
			b.Velocity.X *= 0.99
		}
	}

	if b.Position.Y >= float32(rl.GetScreenHeight()-50) || b.Position.Y <= 0 {
		b.Velocity.Y = -b.Velocity.Y
		return
	}

	b.Sat = float32(int(b.Sat+1) % 360)
}

func (b *Ball) Render() {
	rl.DrawTextureEx(utils.IconTexture, b.Position, 0, 0.05, rl.ColorFromHSV(b.Sat, 1, 1))
}
