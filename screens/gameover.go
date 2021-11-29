package screens

import (
	"fmt"
	"pong/utils"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameOverScreen struct {
	ease_curr float32
	a         float32
	sel       int
	deathMsg  string
}

func (g *GameOverScreen) Reset() {
	g.ease_curr = 0
	g.a = -30 + (utils.EaseInOutQuad(float32(g.ease_curr/1000)) * 30)
	if strings.ToLower(utils.Save.Username) == "örs" {
		g.deathMsg = "Örs ez így nem lesz jó!"
	} else {
		g.deathMsg = "You ded"
	}
}

func (g *GameOverScreen) Update() {
	if g.ease_curr < 1000 {
		g.ease_curr += 20
		g.a = -30 + (utils.EaseInOutQuad(float32(g.ease_curr/1000)) * 30)
	}

	if (rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter)) && g.a > -5 {
		rl.PlaySound(utils.ClickSound)
		//Not falling for it again
		utils.Score = 0
		if g.sel == 0 {
			SwitchToScene(3)
		} else {
			SwitchToScene(0)
		}
	}

	if rl.IsKeyPressed(rl.KeyUp) && g.sel > 0 {
		g.sel--
	}

	if rl.IsKeyPressed(rl.KeyDown) && g.sel < 1 {
		g.sel++
	}
}

func (g *GameOverScreen) Render() {
	utils.DrawCenterText(g.deathMsg, 130, 30, utils.GrayWithAlpha(g.a, 1))
	utils.DrawCenterText(fmt.Sprintf("Score: %d", utils.Score), 170, 25, utils.GrayWithAlpha(g.a, 1))
	//TODO: draw high score
	/* if is_highscore {
		utils.DrawCenterText("New Highscore!", 220, 25, grayWithAlpha(a, 1))
	}
	utils.DrawCenterText(fmt.Sprintf("High score: %d", highscore), 190, 25, grayWithAlpha(a, 1)) */
	utils.DrawTitleBtn("Play again", 0, g.sel, g.a)
	utils.DrawTitleBtn("Main menu", 1, g.sel, g.a)
}
