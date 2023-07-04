package konstructor

import (
	"github.com/hajimehoshi/ebiten"
)

type Position struct {
	X int
	Y int
	Z int
}

type RenderOptions struct {
	Image  string
	Width  int
	Height int
}

type ObjectType struct {
	ID            string
	RenderOptions RenderOptions
}

type Object struct {
	ID       string
	Type     ObjectType
	Position Position
}

type NPCType struct {
	ID            string
	RenderOptions RenderOptions
}

type NPC struct {
	ID       string
	Type     NPCType
	Position Position
}

type Player struct {
	ID            string
	RenderOptions RenderOptions
	Position      Position
}

type ItemType struct {
	ID            string
	RenderOptions RenderOptions
}

type Item struct {
	ID       string
	Type     ItemType
	Position Position
}

type Playground struct {
	Background string
	Objects    []Object
	NPCs       []NPC
	Items      []Item
}

type Level struct {
	ID          string
	Name        string
	Playgrounds []Playground
}

func (e *Engine) PlaygroundDraw(screen *ebiten.Image) {
	e.Domain.GetLevel(e.KContext.CurrentLevel)
}

func (e *Engine) PlaygroundUpdate() {
	e.Domain.GetLevel(e.KContext.CurrentLevel)
}
