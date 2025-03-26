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

func (p *Player) Update() {
}

func (p *Player) Render() {
	rl.DrawTextureEx(p.Texture, p.Position, 0, 1, rl.White)
}

func (p *Player) Deinit() {
  rl.UnloadTexture(p.Texture)
}
