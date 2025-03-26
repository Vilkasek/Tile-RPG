package main

import (
	"tiled/actors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var player actors.Player

func init() {
	rl.InitWindow(1280, 720, "TileRPG")

	player.Init("./assets/graphics/player.png", 0, 0)
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
