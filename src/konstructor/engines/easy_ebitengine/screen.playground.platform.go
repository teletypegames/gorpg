package easy_ebitengine

import "github.com/hajimehoshi/ebiten"

func (e *Engine) PlaygroundPlatformsDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	for _, platform := range playground.Platforms {
		e.PlaygroundAssetDraw(screen, platform.Type.Render, platform.Position)
	}
}
