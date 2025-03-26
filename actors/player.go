package actors

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Texture  rl.Texture2D
	Position rl.Vector2
}

func (p *Player) Init(path string, x int, y int) {
	p.Texture = rl.LoadTexture(path)
	p.Position = rl.Vector2{X: float32(x), Y: float32(y)}
}

func (p *Player) handle_input() {
	// TODO: Add movement
	if rl.IsKeyPressed(rl.KeyW) {
		p.Position.Y -= 32
	} else if rl.IsKeyPressed(rl.KeyS) {
		p.Position.Y += 32
	} else if rl.IsKeyPressed(rl.KeyA) {
		p.Position.X -= 32
	} else if rl.IsKeyPressed(rl.KeyD) {
		p.Position.X += 32
	}
}

func (p *Player) Update() {
	p.handle_input()
}

func (p *Player) Render() {
	rl.DrawTextureEx(p.Texture, p.Position, 0, 1, rl.White)
}

func (p *Player) Deinit() {
	rl.UnloadTexture(p.Texture)
}
