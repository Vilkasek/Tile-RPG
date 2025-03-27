package main

import (
	"tiled/actors"
	userinterface "tiled/user_interface"
	"tiled/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	player actors.Player
	ui     userinterface.UI

	base_res rl.Vector2
)

func init() {
	base_res = rl.Vector2{X: 1024, Y: 576}

	rl.SetConfigFlags(rl.FlagFullscreenMode | rl.FlagWindowResizable)
	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "TileRPG")

	utils.Set_Scale(float32(rl.GetScreenWidth()) / base_res.X)

	player.Init(
		"./assets/graphics/player.png",
		int(256*utils.Get_Scale()),
		int(64*utils.Get_Scale()),
	)
  ui.Init("./assets/graphics/UI-Background.png", 0, 0)
}

func update() {
	player.Update()
  ui.Update()
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	player.Render()
  ui.Render()

	rl.EndDrawing()
}

func deinit() {
	player.Deinit()
  ui.Deinit()
  
	rl.CloseWindow()
}

func main() {
	for !rl.WindowShouldClose() {
		update()
		render()
	}

	deinit()
}
