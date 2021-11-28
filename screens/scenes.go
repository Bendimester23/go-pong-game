package screens

import (
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	scenes = make(map[int]utils.GameObject)

	/*Game state
	0 - Title screen
	1 - Options screen
	2 - Credits screen
	3 - Game screen
	4 - Game over screen
	*/
	state = 0
)

func LoadScenes() {
	scenes[0] = &TitleScreen{}
	scenes[1] = &OptionsScreen{}
	scenes[2] = &CreditsScreen{}
	scenes[3] = &GameScreen{}
	scenes[4] = &GameOverScreen{}
	scenes[5] = &UsernameScreen{}

	ResetAllScenes()
}

func ResetAllScenes() {
	for i := range scenes {
		scenes[i].Reset()
	}
}

func Render() {
	if scenes[state] != nil {
		scenes[state].Render()
		scenes[state].Update()
	} else {
		rl.DrawText("Error: Scene not found", 10, 10, 10, rl.Red)
	}
}

func SwitchToScene(scene int) {
	scenes[state].Reset()
	state = scene
}
