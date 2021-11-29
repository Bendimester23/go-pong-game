package screens

import (
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type OptionsScreen struct {
	selected  int
	ease_curr float32
	a         float32
}

func (o *OptionsScreen) Reset() {
	o.selected = 0
	o.ease_curr = 0
	o.a = -30 + (utils.EaseInOutQuad(float32(o.ease_curr/1000)) * 30)
}

func (o *OptionsScreen) Update() {
	if o.ease_curr < 1000 {
		o.ease_curr += 20
		o.a = -30 + (utils.EaseInOutQuad(float32(o.ease_curr/1000)) * 30)
	}

	if rl.IsKeyPressed(rl.KeyUp) && o.selected > 0 {
		o.selected--
	}

	if rl.IsKeyPressed(rl.KeyDown) && o.selected < 3 {
		o.selected++
	}

	if (rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeySpace)) && o.a > -5 {
		rl.PlaySound(utils.ClickSound)
		utils.SaveGame()
		switch o.selected {
		case 0:
			SwitchToScene(5)
		case 1:
			SwitchToScene(6)
		case 2:
			SwitchToScene(7)
		case 3:
			SwitchToScene(0)
		}
	}
}

func (o *OptionsScreen) Render() {
	utils.DrawCenterTextDef("Options", 80, 30)

	DrawTitleBtn("Set Username", 0, o.selected, o.a, 0)
	DrawTitleBtn("Set Difficulity", 1, o.selected, o.a, 0)
	DrawTitleBtn("Set Weird Mode", 2, o.selected, o.a, 0)
	DrawTitleBtn("Back", 3, o.selected, o.a, 0)
}
