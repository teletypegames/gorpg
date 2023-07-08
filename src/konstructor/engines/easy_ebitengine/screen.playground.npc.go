package easy_ebitengine

import "github.com/hajimehoshi/ebiten"

func (e *Engine) PlaygroundNPCsDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	for _, npc := range playground.NPCs {
		e.PlaygroundAssetDraw(screen, npc.Type.Render, npc.Position)
	}
}
