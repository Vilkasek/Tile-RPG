package main

import (
	"tiled/actors"
	"tiled/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	player   actors.Player
	base_res rl.Vector2
)

func init() {
	base_res = rl.Vector2{X: 1280, Y: 720}

	rl.SetConfigFlags(rl.FlagFullscreenMode | rl.FlagWindowResizable)
	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "TileRPG")

	utils.Set_Scale(float32(rl.GetScreenWidth()) / base_res.X)

	player.Init("./assets/graphics/player.png", 0, 600)
}

func update() {
	player.Update()
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	player.Render()

	rl.EndDrawing()
}

func deinit() {
	player.Deinit()

	rl.CloseWindow()
}

func main() {
	for !rl.WindowShouldClose() {
		update()
		render()
	}

	deinit()
}
