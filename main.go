package main

import rl "github.com/gen2brain/raylib-go/raylib"

func init() {
	rl.InitWindow(1280, 720, "TileRPG")
}

func update() {
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

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
