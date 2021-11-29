package screens

import (
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DifficulityScreen struct {
	selected  int
	ease_curr float32
	a         float32
}

func (d *DifficulityScreen) Reset() {
	d.selected = utils.Save.Difficulty
	d.ease_curr = 0
	d.a = -30 + (utils.EaseInOutQuad(float32(d.ease_curr/1000)) * 30)
}

func (d *DifficulityScreen) Update() {
	if d.ease_curr < 1000 {
		d.ease_curr += 20
		d.a = -30 + (utils.EaseInOutQuad(float32(d.ease_curr/1000)) * 30)
	}

	if rl.IsKeyPressed(rl.KeyUp) && d.selected > 0 {
		d.selected--
	}

	if rl.IsKeyPressed(rl.KeyDown) && d.selected < 5 {
		d.selected++
	}

	if (rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeySpace)) && d.a > -5 {
		rl.PlaySound(utils.ClickSound)
		switch d.selected {
		case 5:
			SwitchToScene(1)
		default:
			utils.Save.Difficulty = d.selected
			utils.SaveGame()
		}
	}
}

func (d *DifficulityScreen) Render() {
	utils.DrawCenterTextDef("Select Difficulity", 80, 30)

	DrawTitleBtn("Baby", 0, d.selected, d.a, 1)
	DrawTitleBtn("Easy", 1, d.selected, d.a, 1)
	DrawTitleBtn("Medium", 2, d.selected, d.a, 1)
	DrawTitleBtn("Hard", 3, d.selected, d.a, 1)
	DrawTitleBtn("TrueChad", 4, d.selected, d.a, 1)
	DrawTitleBtn("Back", 5, d.selected, d.a, 0)
}

func DrawTitleBtn(text string, id, selected int, offset float32, diff int) {
	if utils.Save.Difficulty == id && diff == 1 {
		selected = id
	} else if utils.Save.WeirdMode == (id == 1) && diff == 2 {
		selected = id
	}
	alpha := uint8(255)
	if offset != 0 {
		alpha = uint8((utils.Clamp(0, 30, offset+30.0-float32((id+6)*2)) / 30.0) * 255)
	}
	if selected == id {
		utils.DrawCenterText(text, int32(150+id*30+int(offset)), 20, rl.NewColor(0, 0, 0, alpha))
		lineHeight := float32(170 + id*30 + int(offset))
		txtWidth := float32(rl.MeasureText(text, 20)) / 2
		rl.DrawLineEx(rl.NewVector2(400-txtWidth, lineHeight), rl.NewVector2(400+txtWidth, lineHeight), 2, rl.NewColor(0, 0, 0, alpha))
	} else {
		utils.DrawCenterText(text, int32(150+id*30+int(offset)), 20, rl.NewColor(130, 130, 130, alpha))
	}
}
