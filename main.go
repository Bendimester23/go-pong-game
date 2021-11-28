package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"

	_ "image/png"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	sat float32 = 0

	paddle_lenght int32 = 50

	/*Game state
	0 - Title screen
	1 - Options screen
	2 - Credits screen
	3 - Game screen
	4 - Game over screen
	*/
	state = 0

	vX, vY float32 = 2.0, 2.0

	pAx, pAy, pBx, pBy int32 = 0, 0, 790, 0

	score int32 = 0

	title_pos = rl.NewVector2(365, 50)

	title_anim float32 = 0.0

	title_anim_speed float32 = 28

	title_sel int = 0

	pos = rl.NewVector2(382, 200)

	title_ease_curr float32 = 0

	scenes *map[string]*GameObject
)

func main() {
	rl.InitWindow(800, 450, "Bendi Pong")

	rl.SetTargetFPS(60)

	LoadAllResources()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		switch state {
		case 0:
			renderTitleScreen()
		case 1:
			renderOptionsScreen()
		case 2:
			renderCreditsScreen()
		case 3:
			renderGameScreen()
		case 4:
			renderGameOverScreen()
		}

		sat++
		sat = float32(int(sat) % 360)

		rl.DrawFPS(20, 20)

		rl.EndDrawing()
	}

	UnloadAllResources()

	rl.CloseWindow()
}

func renderCreditsScreen() {
	var a float32 = 0
	if title_ease_curr < 1000 {
		title_ease_curr += 20
		a = -30 + (easeInOutQuad(float32(title_ease_curr/1000)) * 30)
	}

	drawCenterText("Credits", 80+int32(a), 30, grayWithAlpha(a, 1))

	drawCenterText("Made by Bendi", 120+int32(a), 20, grayWithAlpha(a, 2))
	drawCenterText("Written in Raylib-Go", 140+int32(a), 20, grayWithAlpha(a, 4))
	drawCenterText("https://bendi.cf", 160+int32(a), 20, grayWithAlpha(a, 6))
	drawCenterText("https://raylib.com", 180+int32(a), 20, grayWithAlpha(a, 8))

	drawCenterText("Press Space to return", 220+int32(a), 30, grayWithAlpha(a, 10))
	if rl.IsKeyPressed(rl.KeySpace) {
		rl.PlaySound(ClickSound)
		state = 0
		title_ease_curr = 0
	}
}

func grayWithAlpha(a, offset float32) rl.Color {
	return rl.NewColor(100, 100, 100, getAlpha(a, offset))
}

func getAlpha(a, offset float32) uint8 {
	return uint8((clamp(0, 30, a+30.0-offset) / 30.0) * 256)
}

func renderTitleScreen() {
	var a float32 = 0
	if title_ease_curr < 1000 {
		title_ease_curr += 20
		a = -30 + (easeInOutQuad(float32(title_ease_curr/1000)) * 30)
	}
	//Title
	{
		title_anim += title_anim_speed
		if title_anim > 1000 || title_anim < 0 {
			title_anim_speed = -title_anim_speed
		}

		title_pos.Y = 50 + easeInOutQuad(float32(title_anim/1000))*20 + float32(a)

		i_color := rl.ColorFromHSV(sat, 1, 1)

		i_color.A = getAlpha(a, 2)

		rl.DrawTextureEx(IconTexture, title_pos, 0, 0.1, i_color)

		drawCenterText("Bendi Pong", 180+int32(a), 30, rl.NewColor(130, 130, 130, getAlpha(a, 4)))
	}

	//Buttons
	drawTitleBtn("Play", 0, title_sel, a)
	drawTitleBtn("Options", 1, title_sel, a)
	drawTitleBtn("Credits", 2, title_sel, a)
	drawTitleBtn("Quit", 3, title_sel, a)

	if rl.IsKeyPressed(rl.KeyUp) && title_sel > 0 {
		title_sel--
	}

	if rl.IsKeyPressed(rl.KeyDown) && title_sel < 3 {
		title_sel++
	}

	if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeySpace) {
		rl.PlaySound(ClickSound)
		title_ease_curr = 0
		switch title_sel {
		case 0:
			state = 3
		case 1:
			state = 1
		case 2:
			state = 2
		case 3:
			rl.CloseWindow()
			os.Exit(0)
		}
	}
}

func renderOptionsScreen() {
	drawCenterTextDef("Options", 80, 30)

	drawCenterTextDef("WIP", 120, 20)

	drawCenterTextDef("Press Space to return", 200, 20)
	if rl.IsKeyPressed(rl.KeySpace) {
		state = 0
	}
}

var (
	game_over_sel int = 0
)

func renderGameOverScreen() {
	var a float32 = 0
	if title_ease_curr < 1000 {
		title_ease_curr += 20
		a = -30 + (easeInOutQuad(float32(title_ease_curr/1000)) * 30)
	}

	drawCenterText("You ded", 130, 30, grayWithAlpha(a, 1))
	drawCenterText(fmt.Sprintf("Score: %d", score), 170, 25, grayWithAlpha(a, 1))
	if is_highscore {
		drawCenterText("New Highscore!", 220, 25, grayWithAlpha(a, 1))
	}
	drawCenterText(fmt.Sprintf("High score: %d", highscore), 190, 25, grayWithAlpha(a, 1))
	drawTitleBtn("Play again", 0, game_over_sel, a)
	drawTitleBtn("Main menu", 1, game_over_sel, a)
	vX, vY = 2.0, 2.0
	pAx, pAy, pBx, pBy = 0, 0, 790, 0
	pos = rl.NewVector2(40, 40)

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter) {
		rl.PlaySound(ClickSound)
		score = 0
		if game_over_sel == 0 {
			state = 3
		} else {
			state = 0
		}
		title_ease_curr = 0
		return
	}

	if rl.IsKeyPressed(rl.KeyUp) && game_over_sel > 0 {
		game_over_sel--
	}

	if rl.IsKeyPressed(rl.KeyDown) && game_over_sel < 1 {
		game_over_sel++
	}
}

func renderGameScreen() {
	rl.DrawTextureEx(IconTexture, pos, 0, 0.05, rl.ColorFromHSV(sat, 1, 1))

	rl.DrawRectangle(pAx, pAy, 10, paddle_lenght, rl.Red)
	rl.DrawRectangle(pBx, pBy, 10, paddle_lenght, rl.Red)

	rl.DrawText(fmt.Sprintf("Score: %d", score), 5, 420, 30, rl.Gray)

	pos.X += vX
	pos.Y += vY

	if pos.Y >= float32(rl.GetScreenHeight()-50) || pos.Y <= 0 {
		vY = -vY
	}

	if pos.X >= float32(rl.GetScreenWidth()-40) || pos.X <= 10 {
		if ((pAy > int32(pos.Y)+50 || pAy+paddle_lenght < int32(pos.Y)) && pos.X <= 10) ||
			((pBy > int32(pos.Y)+50 || pBy+paddle_lenght < int32(pos.Y)) && pos.X >= float32(rl.GetScreenWidth())-45) {
			state = 4
			rl.PlaySound(WastedSound)
			sendScore()
		} else {
			vX = -vX
			vY = vY + float32(0.05*math.Round(float64(vY/2)))
			vX = vX + float32(0.05*math.Round(float64(vX/2)))
			score++
			rl.PlaySound(ClickSound)
		}
	}

	if rl.IsKeyDown(rl.KeyW) && pAy > 0 {
		pAy -= int32(math.Abs(float64(vX * 5)))
	} else if rl.IsKeyDown(rl.KeyS) && pAy < 400 {
		pAy += int32(math.Abs(float64(vX * 5)))
	}

	if rl.IsKeyDown(rl.KeyUp) && pBy > 0 {
		pBy -= int32(math.Abs(float64(vX * 5)))
	} else if rl.IsKeyDown(rl.KeyDown) && pBy < 400 {
		pBy += int32(math.Abs(float64(vX * 5)))
	}
}

var (
	highscore    = int64(0)
	is_highscore = false
)

func sendScore() {
	res, err := http.Get(fmt.Sprintf("http://31.46.62.84/new?score=%d", score))
	if err != nil {
		panic(err)
	}

	s, _ := ioutil.ReadAll(res.Body)

	if string(s) == "new" {
		is_highscore = true
	}
	fetchHighScore()
}

func fetchHighScore() {
	res, err := http.Get("http://31.46.62.84/best")
	if err != nil {
		panic(err)
	}

	s, _ := ioutil.ReadAll(res.Body)

	highscore, _ = strconv.ParseInt(string(s), 10, 64)
}

func drawTitleBtn(text string, id, selected int, offset float32) {
	alpha := uint8((clamp(0, 30, offset+30.0-float32((id+6)*2)) / 30.0) * 256)
	if selected == id {
		drawCenterText(text, int32(250+id*30+int(offset)), 20, rl.NewColor(0, 0, 0, alpha))
		rl.DrawLineEx(rl.NewVector2(float32(rl.GetScreenWidth()/2-int(rl.MeasureText(text, 20))/2), float32(270+id*30+int(offset))), rl.NewVector2(float32(rl.GetScreenWidth()/2+int(rl.MeasureText(text, 20))/2), float32(270+id*30+int(offset))), 2, rl.NewColor(0, 0, 0, alpha))
	} else {
		drawCenterText(text, int32(250+id*30+int(offset)), 20, rl.NewColor(130, 130, 130, alpha))
	}
}

func clamp(min, max, v float32) float32 {
	if v < min {
		return min
	} else if v > max {
		return max
	}
	return v
}

func drawCenterText(text string, y int32, size int32, color rl.Color) {
	rl.DrawText(text, int32(rl.GetScreenWidth()/2-int(rl.MeasureText(text, size))/2), y, size, color)
}

func drawCenterTextDef(text string, y int32, size int32) {
	rl.DrawText(text, int32(rl.GetScreenWidth()/2-int(rl.MeasureText(text, size))/2), y, size, rl.Gray)
}

//Source: https://easings.net/#easeInOutQuad
func easeInOutQuad(x float32) float32 {
	if x < 0.5 {
		return 2 * x * x
	} else {
		return 1 - (-2*x+2)*(-2*x+2)/2
	}
}
