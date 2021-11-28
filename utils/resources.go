package utils

import (
	"embed"
	"image"
	"io/ioutil"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	resources embed.FS

	WastedSound rl.Sound
	ClickSound  rl.Sound
	IconTexture rl.Texture2D
)

func LoadAllResources(fs embed.FS) {
	resources = fs
	rl.InitAudioDevice()

	IconTexture = loadTexture("resources/texture.png")
	WastedSound = loadSound("resources/wasted.wav")
	ClickSound = loadSound("resources/click.wav")
}

func UnloadAllResources() {
	rl.UnloadSound(WastedSound)
	rl.UnloadSound(ClickSound)
	rl.UnloadTexture(IconTexture)
	rl.CloseAudioDevice()
}

func loadTexture(path string) rl.Texture2D {
	f, err := resources.Open(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	img, _, _ := image.Decode(f)

	return rl.LoadTextureFromImage(rl.NewImageFromImage(img))
}

func loadSound(path string) rl.Sound {
	f, err := resources.Open(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	w := rl.LoadWaveFromMemory(".wav", b, int32(len(b)))

	return rl.LoadSoundFromWave(w)
}
