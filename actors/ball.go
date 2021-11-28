package actors

import (
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Position rl.Vector2
	Velocity utils.Vector2f
	Sat      float32
	LeftBar  *Bar
	RightBar *Bar
}

func (b *Ball) Reset() {
	b.Position = rl.NewVector2(400, 300)
	b.Velocity = utils.Vector2f{X: 1, Y: 1}
	b.Sat = 0
}

func (b *Ball) Update() {
	b.Position.X += b.Velocity.X
	b.Position.Y += b.Velocity.Y

	if b.Position.Y >= float32(rl.GetScreenHeight()-50) || b.Position.Y <= 0 {
		b.Velocity.Y = -b.Velocity.Y
		return
	}
}

func (b *Ball) Render() {
	rl.DrawTextureEx(utils.IconTexture, b.Position, 0, 0.05, rl.ColorFromHSV(b.Sat, 1, 1))
}
