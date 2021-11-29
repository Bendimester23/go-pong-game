package screens

import (
	"os"
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TitleScreen struct {
	selected   int
	ease_curr  float32
	a          float32
	anim       float32
	anim_speed float32
	pos        rl.Vector2

	sat float32
}

func (b *TitleScreen) Reset() {
	b.selected = 0
	b.ease_curr = 0
	b.a = -30 + (utils.EaseInOutQuad(float32(b.ease_curr/1000)) * 30)
	b.anim = 0
	b.anim_speed = 38
	b.pos = rl.NewVector2(365, 50)
	b.sat = 0
}

func (b *TitleScreen) Update() {
	if b.ease_curr < 1000 {
		b.ease_curr += 20
		b.a = -30 + (utils.EaseInOutQuad(float32(b.ease_curr/1000)) * 30)
	}

	b.sat++
	b.sat = float32(int(b.sat) % 360)

	b.anim += b.anim_speed
	if b.anim > 1000 || b.anim < 0 {
		b.anim_speed = -b.anim_speed
	}

	b.pos.Y = 50 + utils.EaseInOutQuad(float32(b.anim/1000))*20 + float32(b.a)

	if rl.IsKeyPressed(rl.KeyUp) && b.selected > 0 {
		b.selected--
	}

	if rl.IsKeyPressed(rl.KeyDown) && b.selected < 3 {
		b.selected++
	}

	if (rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeySpace)) && b.a > -5 {
		rl.PlaySound(utils.ClickSound)
		switch b.selected {
		case 0:
			SwitchToScene(3)
		case 1:
			SwitchToScene(1)
		case 2:
			SwitchToScene(2)
		case 3:
			utils.SaveGame()
			rl.CloseWindow()
			os.Exit(0)
		}
	}
}

func (b *TitleScreen) Render() {
	//Title
	i_color := rl.ColorFromHSV(b.sat, 1, 1)

	i_color.A = utils.GetAlpha(b.a, 2)

	rl.DrawTextureEx(utils.IconTexture, b.pos, 0, 0.1, i_color)

	utils.DrawCenterText("Bendi Pong", 180+int32(b.a), 30, rl.NewColor(130, 130, 130, utils.GetAlpha(b.a, 4)))

	//Buttons
	utils.DrawTitleBtn("Play", 0, b.selected, b.a)
	utils.DrawTitleBtn("Options", 1, b.selected, b.a)
	utils.DrawTitleBtn("Credits", 2, b.selected, b.a)
	utils.DrawTitleBtn("Quit", 3, b.selected, b.a)
}
