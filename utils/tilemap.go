package utils

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tilemap struct {
	Tileset rl.Texture2D
	Tile    rl.Rectangle

	Grid [32][18]int
}

func (m *Tilemap) setup_grid() {
	for i := 0; i < 32; i++ {
		for j := 0; j < 18; j++ {
			m.Grid[i][j] = rand.Intn(2)
		}
	}
}

func (m *Tilemap) Init(path string) {
	m.setup_grid()

	m.Tileset = rl.LoadTexture(path)
}

func (m *Tilemap) Render() {
	var destRecs []rl.Rectangle
	var sourceRecs []rl.Rectangle

	for i := 0; i < 32; i++ {
		for j := 0; j < 18; j++ {
			switch value := m.Grid[i][j]; value {
			case 1:
				m.Tile = rl.NewRectangle(32, 0, 32, 32)
			case 0:
				m.Tile = rl.NewRectangle(0, 0, 32, 32)
			}

			sourceRecs = append(sourceRecs, m.Tile)

			destRecs = append(destRecs,
				rl.NewRectangle(
					(float32(i)*32)*Get_Scale(),
					(float32(j)*32)*Get_Scale(),
					32*Get_Scale(),
					32*Get_Scale(),
				),
			)
		}
	}
	origin := rl.NewVector2(0, 0)

	rotation := float32(0)
	if len(destRecs) > 0 {
		for index, destRec := range destRecs {
			rl.DrawTexturePro(m.Tileset, sourceRecs[index], destRec, origin, rotation, rl.White)
		}
	}
}
