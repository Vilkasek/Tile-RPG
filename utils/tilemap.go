package utils

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tilemap struct {
	Tiles    [][]int
	TileSize int32
	Texture  rl.Texture2D
	Width    int
	Height   int
}

func NewTilemap(width, height int, tileSize int32, path string) *Tilemap {
	tiles := make([][]int, height)
	for y := range tiles {
		tiles[y] = make([]int, width)
		for x := range tiles[y] {
			tiles[y][x] = rand.Intn(4)
		}
	}

	texture := rl.LoadTexture(path)

	return &Tilemap{
		Tiles:    tiles,
		TileSize: tileSize,
		Width:    width,
		Height:   height,
		Texture:  texture,
	}
}

func (tm *Tilemap) Render() {
	for y, row := range tm.Tiles {
		for x, tile := range row {
			source := rl.Rectangle{
				X:      float32(tile * int(tm.TileSize)),
				Y:      0,
				Width:  float32(tm.TileSize),
				Height: float32(tm.TileSize),
			}
			dest := rl.Rectangle{
				X:      float32(x*int(tm.TileSize)) * scale,
				Y:      float32(y*int(tm.TileSize)) * scale,
				Width:  float32(tm.TileSize) * scale,
				Height: float32(tm.TileSize) * scale,
			}
			rl.DrawTexturePro(tm.Texture, source, dest, rl.Vector2{}, 0, rl.White)
		}
	}
}
