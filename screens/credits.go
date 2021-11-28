package screens

import (
	"pong/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CreditsScreen struct {
	ease_curr float32
	a         float32
}

func (c *CreditsScreen) Reset() {
	c.ease_curr = 0
	c.a = -30 + (utils.EaseInOutQuad(float32(c.ease_curr/1000)) * 30)
}

func (c *CreditsScreen) Update() {
	if c.ease_curr < 1000 {
		c.ease_curr += 20
		c.a = -30 + (utils.EaseInOutQuad(float32(c.ease_curr/1000)) * 30)
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		rl.PlaySound(utils.ClickSound)
		SwitchToScene(0)
	}
}

func (c *CreditsScreen) Render() {
	utils.DrawCenterText("Credits", 80+int32(c.a), 30, utils.GrayWithAlpha(c.a, 1))

	utils.DrawCenterText("Made by Bendi", 120+int32(c.a), 20, utils.GrayWithAlpha(c.a, 2))
	utils.DrawCenterText("Written in Raylib-Go", 140+int32(c.a), 20, utils.GrayWithAlpha(c.a, 4))
	utils.DrawCenterText("https://bendi.cf", 160+int32(c.a), 20, utils.GrayWithAlpha(c.a, 6))
	utils.DrawCenterText("https://raylib.com", 180+int32(c.a), 20, utils.GrayWithAlpha(c.a, 8))

	utils.DrawCenterText("Press Space to return", 220+int32(c.a), 30, utils.GrayWithAlpha(c.a, 10))
}
