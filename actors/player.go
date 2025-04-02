package actors

import (
	"tiled/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Texture  rl.Texture2D
	Position rl.Vector2
}

func (p *Player) Init(path string, x int, y int) {
	p.Texture = rl.LoadTexture(path)
	p.Position = rl.Vector2{X: float32(x), Y: float32(y)}
}

func (p *Player) clamp_player() {
	if p.Position.Y <= 64*utils.Get_Scale() {
		p.Position.Y = 64 * utils.Get_Scale()
	}
	if p.Position.Y+float32(p.Texture.Height)*utils.Get_Scale() >= float32(rl.GetScreenHeight()) {
		p.Position.Y = float32(rl.GetScreenHeight()) - float32(p.Texture.Height)*utils.Get_Scale()
	}

	if p.Position.X <= 0 {
		p.Position.X = 0
	}
	if p.Position.X+float32(p.Texture.Width)*utils.Get_Scale() >= float32(rl.GetScreenWidth()) {
		p.Position.X = float32(rl.GetScreenWidth()) - float32(p.Texture.Width)*utils.Get_Scale()
	}
}

// func (p *Player) collect_resources(t *utils.Tilemap) {
// 	if rl.IsKeyPressed(rl.KeySpace) {
// 		tileX := int(p.Position.X / (32 * utils.Get_Scale()))
// 		tileY := int(p.Position.Y / (32 * utils.Get_Scale()))
//
// 		switch value := t.Get_Tile_Value(tileX, tileY); value {
// 		case 1:
// 			utils.Add_Fiber(5)
// 		case 2:
// 			utils.Add_Stone(5)
// 		case 3:
// 			utils.Add_Wood(5)
// 		}
//
// 		t.Set_Tile_Value(tileX, tileY, 0)
// 	}
// }

func (p *Player) handle_input( /*t *utils.Tilemap*/ ) {
	if rl.IsKeyPressed(rl.KeyW) {
		p.Position.Y -= 32 * utils.Get_Scale()
	} else if rl.IsKeyPressed(rl.KeyS) {
		p.Position.Y += 32 * utils.Get_Scale()
	} else if rl.IsKeyPressed(rl.KeyA) {
		p.Position.X -= 32 * utils.Get_Scale()
	} else if rl.IsKeyPressed(rl.KeyD) {
		p.Position.X += 32 * utils.Get_Scale()
	}

	p.clamp_player()

	// p.collect_resources(t)
}

func (p *Player) Update(t *utils.Tilemap) {
	p.handle_input()
}

func (p *Player) Render() {
	rl.DrawTextureEx(p.Texture, p.Position, 0, utils.Get_Scale(), rl.White)
}

func (p *Player) Deinit() {
	rl.UnloadTexture(p.Texture)
}
