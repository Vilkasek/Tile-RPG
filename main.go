package main

import (
	"tiled/actors"
	userinterface "tiled/user_interface"
	"tiled/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	player  actors.Player
	ui      userinterface.UI
	tilemap utils.Tilemap

	base_res rl.Vector2
)

func init() {
	base_res = rl.Vector2{X: 1024, Y: 576}

	rl.SetConfigFlags(rl.FlagFullscreenMode | rl.FlagWindowResizable)
	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "TileRPG")

	utils.Set_Scale(float32(rl.GetScreenWidth()) / base_res.X)

	player.Init(
		"./assets/graphics/Player.png",
		int(512*utils.Get_Scale()),
		int(288*utils.Get_Scale()),
	)
	ui.Init("./assets/graphics/UI-Background.png", "./assets/graphics/Icons.png", 0, 0)
	tilemap.Init("./assets/graphics/Tiles.png")
}

func update() {
	player.Update(&tilemap)
	ui.Update()
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	tilemap.Render()
	ui.Render()
	player.Render()

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
