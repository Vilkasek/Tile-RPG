package userinterface

import (
	"tiled/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UI struct {
	Texture rl.Texture2D
	icon    rl.Texture2D

	icons []rl.Rectangle

	Position rl.Vector2
}

func (u *UI) Init(path, icon string, x int, y int) {
	u.Texture = rl.LoadTexture(path)
	u.icon = rl.LoadTexture(icon)

	u.icons = append(
		u.icons,
		rl.NewRectangle(0, 0, 64, 64),
		rl.NewRectangle(64, 0, 64, 64),
		rl.NewRectangle(128, 0, 64, 64),
	)

	u.Position = rl.Vector2{X: float32(x), Y: float32(y)}
}

func (u *UI) Deinit() {
	rl.UnloadTexture(u.Texture)
}

func (u *UI) Update() {
}

func (u *UI) Render() {
	sourceRecs := u.icons

	var destRecs []rl.Rectangle

	for i := 0; i < 3; i++ {
		destRecs = append(
			destRecs,
			rl.NewRectangle(
				float32(i*64),
				0,
				64*utils.Get_Scale(),
				64*utils.Get_Scale(),
			),
		)
	}

	rl.DrawTextureEx(u.Texture, u.Position, 0, utils.Get_Scale(), rl.White)

	if len(destRecs) > 0 {
		for index, destRec := range destRecs {
			rl.DrawTexturePro(u.icon, sourceRecs[index], destRec, rl.Vector2Zero(), 0, rl.White)
		}
	}
}
