package utils

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MapGenerator struct {
	Tileset rl.Texture2D
	Tile    rl.Rectangle
}

func (m *MapGenerator) Init(path string) {
	m.Tileset = rl.LoadTexture(path)
	m.Tile = rl.NewRectangle(0, 0, 32, 32)
}

func (m *MapGenerator) Render() {
	sourceRec := m.Tile
	destRec := rl.NewRectangle(0, 64*Get_Scale(), 32*Get_Scale(), 32*Get_Scale())

	origin := rl.NewVector2(0, 0)

	rotation := float32(0)

	rl.DrawTexturePro(m.Tileset, sourceRec, destRec, origin, rotation, rl.White)
}
