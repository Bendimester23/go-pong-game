package screens

import (
	"fmt"
	"pong/actors"
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScreen struct {
	ball     *actors.Ball
	leftBar  *actors.Bar
	rightBar *actors.Bar
}

func (b *GameScreen) Reset() {
	b.leftBar = &actors.Bar{
		X:       0,
		Height:  50,
		Width:   10,
		Speed:   10,
		KeyUp:   rl.KeyW,
		KeyDown: rl.KeyS,
		Color:   rl.Red,
	}

	b.rightBar = &actors.Bar{
		X:       790,
		Height:  50,
		Width:   10,
		Speed:   10,
		KeyUp:   rl.KeyUp,
		KeyDown: rl.KeyDown,
		Color:   rl.Red,
	}

	b.ball = &actors.Ball{
		LeftBar:  b.leftBar,
		RightBar: b.rightBar,
	}

	b.ball.Reset()
	b.leftBar.Reset()
	b.rightBar.Reset()
}

func (g *GameScreen) Update() {
	g.leftBar.Update()
	g.rightBar.Update()
	g.ball.Update()
	b := g.ball
	if b.Position.X <= 10 {
		if b.LeftBar.Y > int32(b.Position.Y)+50 || b.LeftBar.Y+b.LeftBar.Height < int32(b.Position.Y) {
			rl.PlaySound(utils.WastedSound)
			SwitchToScene(4)
			return
		}
		utils.Score++
		b.Velocity.X = -b.Velocity.X
		rl.PlaySound(utils.ClickSound)
	}
	if b.Position.X >= float32(rl.GetScreenWidth()-40) {
		if b.RightBar.Y > int32(b.Position.Y)+50 || b.RightBar.Y+b.RightBar.Height < int32(b.Position.Y) {
			rl.PlaySound(utils.WastedSound)
			SwitchToScene(4)
			utils.SaveGame()
			return
		}
		utils.Score++
		rl.PlaySound(utils.ClickSound)
		b.Velocity.X = -b.Velocity.X
	}
}

func (b *GameScreen) Render() {
	b.leftBar.Render()
	b.rightBar.Render()
	b.ball.Render()
	rl.DrawText(fmt.Sprintf("Score: %d", utils.Score), 5, 420, 30, rl.Gray)
}
