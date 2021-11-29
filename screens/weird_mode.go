package screens

import (
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WeirdModeScreen struct {
	selected  int
	ease_curr float32
	a         float32
}

func (d *WeirdModeScreen) Reset() {
	d.selected = utils.Save.Difficulty
	d.ease_curr = 0
	d.a = -30 + (utils.EaseInOutQuad(float32(d.ease_curr/1000)) * 30)
}

func (d *WeirdModeScreen) Update() {
	if d.ease_curr < 1000 {
		d.ease_curr += 20
		d.a = -30 + (utils.EaseInOutQuad(float32(d.ease_curr/1000)) * 30)
	}

	if rl.IsKeyPressed(rl.KeyUp) && d.selected > 0 {
		d.selected--
	}

	if rl.IsKeyPressed(rl.KeyDown) && d.selected < 2 {
		d.selected++
	}

	if (rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeySpace)) && d.a > -5 {
		rl.PlaySound(utils.ClickSound)
		switch d.selected {
		case 2:
			SwitchToScene(1)
		default:
			utils.Save.WeirdMode = d.selected == 1
			utils.SaveGame()
		}
	}
}

func (d *WeirdModeScreen) Render() {
	utils.DrawCenterTextDef("Weird Mode", 80, 30)

	DrawTitleBtn("Off", 0, d.selected, d.a, 2)
	DrawTitleBtn("On", 1, d.selected, d.a, 2)
	DrawTitleBtn("Back", 2, d.selected, d.a, 0)
}
