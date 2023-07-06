package domain

import "game/konstructor"

func (d *Domain) InitLevel() {
	d.Levels = []konstructor.Level{
		{
			ID:   "level_1",
			Name: "Level I.",
			Playgrounds: []konstructor.Playground{
				{
					Render: konstructor.Render{
						Image: "assets/images/playgrounds/level_1_playground_1.png",
					},
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
								Render: konstructor.Render{
									Image:  "assets/images/objects/test_object.png",
									Width:  30,
									Height: 30,
								},
							},
						},
						{
							ID: "test_object2",
							Position: konstructor.Position{
								X: 50,
								Y: 50,
								Z: 0,
							},
							Type: konstructor.ObjectType{
								ID: "test_object_type",
								Render: konstructor.Render{
									Image:  "assets/images/objects/test_object.png",
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

func (d *Domain) GetLevels() []konstructor.Level {
	return d.Levels
}

func (d *Domain) GetLevel(index int) konstructor.Level {
	return d.Levels[index]
}
