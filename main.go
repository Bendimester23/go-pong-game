package main

import (
	"embed"
	"flag"

	_ "image/png"
	"pong/screens"
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//go:embed resources
	res embed.FS

	showFps = flag.Bool("fps", false, "show fps")
)

func main() {
	rl.InitWindow(800, 450, "Bendi Pong")

	rl.SetTargetFPS(60)

	utils.LoadSave()

	utils.LoadAllResources(res)
	screens.LoadScenes()

	if !utils.Save.Saved {
		screens.SwitchToScene(5)
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		screens.Render()

		if *showFps {
			rl.DrawFPS(20, 20)
		}

		rl.EndDrawing()
	}

	utils.UnloadAllResources()

	rl.CloseWindow()
}
