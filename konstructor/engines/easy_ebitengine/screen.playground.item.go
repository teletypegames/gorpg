package easy_ebitengine

import "github.com/hajimehoshi/ebiten"

func (e *Engine) PlaygroundItemsDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	for _, item := range playground.Items {
		e.PlaygroundAssetDraw(screen, item.Type.Render, item.Position)
	}
}
