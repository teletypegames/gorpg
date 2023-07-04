package domain

import "game/konstructor/entity"

func (d *Domain) InitLevel() {
	d.Levels = []entity.Level{
		{
			ID:   "level_1",
			Name: "Level I.",
			Playgrounds: []entity.Playground{
				{
					Background: "level_1_playground_1.png",
					Objects: []entity.Object{
						{
							ID: "test_object",
							Position: entity.Position{
								X: 10,
								Y: 10,
								Z: 0,
							},
							Type: entity.ObjectType{
								ID: "test_object_type",
								RenderOptions: entity.RenderOptions{
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

func (d *Domain) GetLevel(index int) entity.Level {
	return d.Levels[index]
}
