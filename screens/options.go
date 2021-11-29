package screens

import (
	"os"
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

	if rl.IsKeyPressed(rl.KeyDown) && o.selected < 1 {
		o.selected++
	}

	if (rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeySpace)) && o.a > -5 {
		rl.PlaySound(utils.ClickSound)
		utils.SaveGame()
		switch o.selected {
		case 0:
			SwitchToScene(5)
		case 1:
			SwitchToScene(0)
		case 2:
			SwitchToScene(2)
		case 3:
			utils.SaveGame()
			rl.CloseWindow()
			os.Exit(0)
		}
	}
}

func (o *OptionsScreen) Render() {
	utils.DrawCenterTextDef("Options", 80, 30)

	utils.DrawTitleBtn("Set Username", 0, o.selected, o.a)
	utils.DrawTitleBtn("Back", 1, o.selected, o.a)
}
