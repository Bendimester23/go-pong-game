package screens

import (
	"pong/utils"
	"regexp"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UsernameScreen struct {
	val   string
	valid bool
}

func (o *UsernameScreen) Reset() {
	o.val = ""
}

func (o *UsernameScreen) Update() {
	if rl.IsKeyPressed(rl.KeyBackspace) {
		if len(o.val) > 0 {
			o.val = o.val[:len(o.val)-1]
		}
	}

	c := rl.GetCharPressed()
	if c != 0 && len(o.val) < 16 {
		o.val += string(c)
	}

	if b, _ := regexp.MatchString("^[a-zA-Z0-9_áÁéÉöÖüÜóÓőŐúÚűŰäÄß]{3,16}$", o.val); b {
		o.valid = true
	} else {
		o.valid = false
	}

	if (rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter)) && o.valid {
		utils.Save.Username = o.val
		utils.SaveGame()
		SwitchToScene(0)
	}
}

func (o *UsernameScreen) Render() {
	utils.DrawCenterTextDef("Enter your username", 80, 30)

	l := float32(rl.MeasureText(o.val, 20)/2) + 10

	if o.valid {
		utils.DrawCenterTextDef(o.val, 130, 20)
		rl.DrawLineEx(rl.NewVector2(400-l, 152), rl.NewVector2(400+l, 152), 2, rl.Gray)
	} else {
		utils.DrawCenterText(o.val, 130, 20, rl.Red)
		rl.DrawLineEx(rl.NewVector2(400-l, 152), rl.NewVector2(400+l, 152), 2, rl.Red)
		utils.DrawCenterText("Invalid username", 160, 16, rl.Red)
	}

	utils.DrawCenterTextDef("Press Space or Return to Continue", 200, 20)
}
