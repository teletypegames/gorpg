package easy_ebitengine

import "github.com/hajimehoshi/ebiten"

func (e *Engine) PlaygroundDefaultPlayerDraw(screen *ebiten.Image) {
	player := e.Domain.GetDefaultPlayer()
	e.PlaygroundAssetDraw(screen, player.Type.Render, player.Position)
}
