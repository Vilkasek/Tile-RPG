package main

import rl "github.com/gen2brain/raylib-go/raylib"

var player rl.Texture2D

func init() {
	rl.InitWindow(1280, 720, "TileRPG")

	player = rl.LoadTexture("./assets/graphics/player.png")
}

func update() {
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	rl.DrawTexture(player, 100, 100, rl.White)

	rl.EndDrawing()
}

func deinit() {
	rl.CloseWindow()
}

func main() {
	for !rl.WindowShouldClose() {
		update()
		render()
	}

	deinit()
}
