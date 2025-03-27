package userinterface

import (
	"tiled/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UI struct {
	Texture  rl.Texture2D
	Position rl.Vector2
}

func (u *UI) Init(path, icon string, x int, y int) {
	u.Texture = rl.LoadTexture(path)
	u.Position = rl.Vector2{X: float32(x), Y: float32(y)}
}

func (u *UI) Deinit() {
	rl.UnloadTexture(u.Texture)
}

func (u *UI) Update() {
}

func (u *UI) Render() {
	rl.DrawTextureEx(u.Texture, u.Position, 0, utils.Get_Scale(), rl.White)
}
