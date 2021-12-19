package screens

import (
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LeaderboardScreen struct {
	ease_curr float32
	a         float32
	sel       int32
}

func (l *LeaderboardScreen) Reset() {
	l.ease_curr = 0
	l.a = -30 + (utils.EaseInOutQuad(float32(l.ease_curr/1000)) * 30)
}

func (l *LeaderboardScreen) Update() {
	if l.ease_curr < 1000 {
		l.ease_curr += 20
		l.a = -30 + (utils.EaseInOutQuad(float32(l.ease_curr/1000)) * 30)
	}

	if rl.IsKeyPressed(rl.KeyUp) {
		if l.sel > 0 {
			l.sel--
		}
	}

	if rl.IsKeyPressed(rl.KeyDown) {
		if l.sel < 10 {
			l.sel++
		}
	}
}

func (l *LeaderboardScreen) Render() {
	utils.DrawCenterText("Leaderboard", 80+int32(l.a), 30, utils.GrayWithAlpha(l.a, 1))

	drawEntry(" 1. Player - 6969 points", 0, l.a, l.sel)
	drawEntry(" 2. Player - 6969 points", 1, l.a, l.sel)
	drawEntry(" 3. Player - 6969 points", 2, l.a, l.sel)
	drawEntry(" 4. Player - 6969 points", 3, l.a, l.sel)
	drawEntry(" 5. Player - 6969 points", 4, l.a, l.sel)
	drawEntry(" 6. Player - 6969 points", 5, l.a, l.sel)
	drawEntry(" 7. Player - 6969 points", 6, l.a, l.sel)
	drawEntry(" 8. Player - 6969 points", 7, l.a, l.sel)
	drawEntry(" 9. Player - 6969 points", 8, l.a, l.sel)
	drawEntry("10. Player - 6969 points", 9, l.a, l.sel)
	drawEntry("6969.  You - -69  points", 10, l.a, l.sel)
}

func drawEntry(txt string, id int32, a float32, sel int32) {
	h := float32(120+(id*25)+20) + a
	utils.DrawCenterText(txt, 120+(id*25)+int32(a), 20, utils.GrayWithAlpha(a, float32(id+1)))
	if sel == id {
		l := float32(rl.MeasureText(txt, 20) / 2)
		rl.DrawLineEx(rl.Vector2{X: 400 - l, Y: h}, rl.Vector2{X: 400 + l, Y: h}, 2, utils.GrayWithAlpha(a, float32(id+1)))
	}
}
