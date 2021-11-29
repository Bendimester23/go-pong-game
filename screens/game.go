package screens

import (
	"fmt"
	"math/rand"
	"pong/actors"
	"pong/utils"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScreen struct {
	ball      *actors.Ball
	leftBar   *actors.Bar
	rightBar  *actors.Bar
	lastHit   time.Time
	countDown int
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
		Speed:    1 + float32(utils.Save.Difficulty)/5,
	}

	b.countDown = 180

	b.lastHit = time.Now()

	b.ball.Reset()
	b.leftBar.Reset()
	b.rightBar.Reset()

	if utils.Save.Difficulty == 0 {
		b.leftBar.Height = 100
		b.rightBar.Height = 100
	}

	if utils.Save.Difficulty == 4 {
		b.leftBar.Height = 25
		b.rightBar.Height = 25
	}
}

func (g *GameScreen) Update() {
	g.leftBar.Update()
	g.rightBar.Update()
	if g.countDown == 1 {
		rl.PlaySound(utils.ClickSound)
	}
	if g.countDown > 0 {
		g.countDown--
	} else {
		g.ball.Update()
	}
	b := g.ball
	if b.Position.X <= 10 && time.Since(g.lastHit).Seconds() > 1 {
		g.lastHit = time.Now()
		if b.LeftBar.Y > int32(b.Position.Y)+50 || b.LeftBar.Y+b.LeftBar.Height < int32(b.Position.Y) {
			rl.PlaySound(utils.WastedSound)
			SwitchToScene(4)
			return
		}
		incrementScore(b)
	}
	if b.Position.X >= float32(rl.GetScreenWidth()-40) && time.Since(g.lastHit).Seconds() > 1 {
		g.lastHit = time.Now()
		if b.RightBar.Y > int32(b.Position.Y)+50 || b.RightBar.Y+b.RightBar.Height < int32(b.Position.Y) {
			rl.PlaySound(utils.WastedSound)
			SwitchToScene(4)
			utils.SaveGame()
			return
		}
		incrementScore(b)
	}
}

func incrementScore(b *actors.Ball) {
	rl.PlaySound(utils.ClickSound)
	utils.Score++
	b.Velocity.X = -b.Velocity.X
	if utils.Save.Difficulty == 0 {
		return
	}
	if utils.Save.Difficulty == 4 {
		b.Velocity.X = b.Velocity.X * (float32(utils.Save.Difficulty+1)/40 + 1) * randomMultiplier()
		b.Velocity.Y = b.Velocity.Y * (float32(utils.Save.Difficulty+1)/40 + 1) * randomMultiplier()
		return
	}
	if utils.Score < 10 {
		b.Velocity.X = b.Velocity.X + float32(utils.Save.Difficulty)/10
		b.Velocity.Y = b.Velocity.Y + float32(utils.Save.Difficulty)/10
	} else if utils.Score < 15 {
		b.Velocity.X = b.Velocity.X + float32(utils.Save.Difficulty)/8
		b.Velocity.Y = b.Velocity.Y + float32(utils.Save.Difficulty)/8
	} else if utils.Score < 20 {
		b.Velocity.X = b.Velocity.X + float32(utils.Save.Difficulty)/6
		b.Velocity.Y = b.Velocity.Y + float32(utils.Save.Difficulty)/6
	} else if utils.Score < 25 {
		b.Velocity.X = b.Velocity.X + float32(utils.Save.Difficulty)/4
		b.Velocity.Y = b.Velocity.Y + float32(utils.Save.Difficulty)/4
	} else if utils.Score < 30 {
		b.Velocity.X = b.Velocity.X + float32(utils.Save.Difficulty)/2
		b.Velocity.Y = b.Velocity.Y + float32(utils.Save.Difficulty)/2
	} else if utils.Score < 35 {
		b.Velocity.X = b.Velocity.X + float32(utils.Save.Difficulty)
		b.Velocity.Y = b.Velocity.Y + float32(utils.Save.Difficulty)
	} else {
		b.Velocity.X = b.Velocity.X * (float32(utils.Save.Difficulty)/10 + 1)
		b.Velocity.Y = b.Velocity.Y * (float32(utils.Save.Difficulty)/10 + 1)
	}
	b.Velocity.X = b.Velocity.X * randomMultiplier()
	b.Velocity.Y = b.Velocity.Y * randomMultiplier()
}

func randomMultiplier() float32 {
	return rand.Float32()/4 + 0.85
}

func (g *GameScreen) Render() {
	g.leftBar.Render()
	g.rightBar.Render()
	if g.countDown > 0 {
		utils.DrawCenterText(fmt.Sprintf("%d", int(g.countDown/60)), int32(rl.GetScreenHeight()/2-18), 36, rl.DarkGray)
	} else {
		g.ball.Render()
	}
	rl.DrawText(fmt.Sprintf("Score: %d", utils.Score), 5, 420, 30, rl.Gray)
}
