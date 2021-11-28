package screens

import (
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type OptionsScreen struct {
}

func (o *OptionsScreen) Reset() {

}

func (o *OptionsScreen) Update() {
	if rl.IsKeyPressed(rl.KeySpace) {
		SwitchToScene(0)
	}
}

func (o *OptionsScreen) Render() {
	utils.DrawCenterTextDef("Options", 80, 30)

	utils.DrawCenterTextDef("WIP", 120, 20)

	utils.DrawCenterTextDef("Press Space to return", 200, 20)
}
