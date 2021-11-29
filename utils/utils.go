package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameObject interface {
	Update()
	Render()
	Reset()
}

type Options struct {
	Username   string `json:"username"`
	Saved      bool   `kson:"_"`
	HighScore  int    `json:"highscore"`
	Difficulty int    `json:"difficulty"`
	WeirdMode  bool   `json:"weird_mode"`
}

var (
	Save *Options
)

func DrawTitleBtn(text string, id, selected int, offset float32) {
	alpha := uint8((Clamp(0, 30, offset+30.0-float32((id+6)*2)) / 30.0) * 256)
	if selected == id {
		DrawCenterText(text, int32(250+id*30+int(offset)), 20, rl.NewColor(0, 0, 0, alpha))
		rl.DrawLineEx(rl.NewVector2(float32(rl.GetScreenWidth()/2-int(rl.MeasureText(text, 20))/2), float32(270+id*30+int(offset))), rl.NewVector2(float32(rl.GetScreenWidth()/2+int(rl.MeasureText(text, 20))/2), float32(270+id*30+int(offset))), 2, rl.NewColor(0, 0, 0, alpha))
	} else {
		DrawCenterText(text, int32(250+id*30+int(offset)), 20, rl.NewColor(130, 130, 130, alpha))
	}
}

func DrawCenterText(text string, y int32, size int32, color rl.Color) {
	rl.DrawText(text, int32(rl.GetScreenWidth()/2-int(rl.MeasureText(text, size))/2), y, size, color)
}

func DrawCenterTextDef(text string, y int32, size int32) {
	rl.DrawText(text, int32(rl.GetScreenWidth()/2-int(rl.MeasureText(text, size))/2), y, size, rl.Gray)
}

func GetDataFile() string {
	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf("%s\\pong-save.json", os.Getenv("APPDATA"))
	case "linux":
		return fmt.Sprintf("%s/pong-save.json", os.Getenv("HOME"))
	}
	a, _ := os.Getwd()
	return a
}

func LoadSave() {
	f, err := os.Open(GetDataFile())
	if err != nil {
		Save = &Options{}
		return
	}
	defer f.Close()

	raw, err := ioutil.ReadAll(f)

	if err != nil {
		Save = &Options{}
		return
	}

	json.Unmarshal(raw, &Save)
}

func SaveGame() {
	f, err := os.Create(GetDataFile())
	if err != nil {
		fmt.Printf("Error saving: %s", err.Error())
		return
	}
	defer f.Close()

	Save.Saved = true

	raw, err := json.Marshal(Save)
	if err != nil {
		fmt.Printf("Error saving: %s", err.Error())
		return
	}

	f.Write(raw)
}
