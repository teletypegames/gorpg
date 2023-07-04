package domain

import (
	"game/konstructor"
)

func (d *Domain) InitLevel() {
	d.Levels = []konstructor.Level{
		{
			ID:   "level_1",
			Name: "Level I.",
			Playgrounds: []konstructor.Playground{
				{
					Background: "level_1_playground_1.png",
					Objects: []konstructor.Object{
						{
							ID: "test_object",
							Position: konstructor.Position{
								X: 10,
								Y: 10,
								Z: 0,
							},
							Type: konstructor.ObjectType{
								ID: "test_object_type",
								RenderOptions: konstructor.RenderOptions{
									Image:  "test_object.png",
									Width:  30,
									Height: 30,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *Domain) GetLevel(index int) konstructor.Level {
	return d.Levels[index]
}
