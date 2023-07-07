package domain

import "game/konstructor"

const (
	Level1Playground1 konstructor.PlaygroundID = "level_1_playground_1"
)

func (d *Domain) InitLevel() {
	d.Levels = []konstructor.Level{
		{
			ID:   "level_1",
			Name: "Level I.",
			Playgrounds: []konstructor.Playground{
				{
					ID: Level1Playground1,
					Render: konstructor.Render{
						Image: konstructor.GetPlaygroundImagePath(Level1Playground1),
					},
					Platforms: []konstructor.Platform{
						{
							ID: "test_platform",
							Position: konstructor.Position{
								X: 10,
								Y: 10,
								Z: 0,
							},
							Type: d.GetPlatformType(TestPlatformType),
						},
						{
							ID: "test_platform2",
							Position: konstructor.Position{
								X: 50,
								Y: 50,
								Z: 0,
							},
							Type: d.GetPlatformType(TestPlatformType),
						},
					},
					Items: []konstructor.Item{
						{
							ID: "sword_1",
							Position: konstructor.Position{
								X: 100,
								Y: 100,
								Z: 0,
							},
							Type: d.GetItemType(SwordItemType),
						},
					},
					NPCs: []konstructor.NPC{
						{
							ID: "test_npc_1",
							Position: konstructor.Position{
								X: 220,
								Y: 180,
								Z: 0,
							},
							Type: d.GetNPCType(TestNPCType),
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
