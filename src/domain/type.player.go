package domain

import "git.teletypegames.org/engines/gorpg/src/konstructor"

const (
	DefaultPlayerType konstructor.PlayerTypeID = "default"
)

func (d *Domain) InitPlayer() {
	d.Players = []konstructor.Player{
		{
			ID: "default",
			Type: konstructor.PlayerType{
				ID: DefaultPlayerType,
				Render: konstructor.Render{
					Image: konstructor.GetPlayerTypeImagePath(DefaultPlayerType),
				},
			},
			Position: konstructor.Position{
				X: 70,
				Y: 70,
				Z: 1,
			},
		},
	}
}

func (d *Domain) GetPlayers() []konstructor.Player {
	return d.Players
}

func (d *Domain) GetPlayer(index int) konstructor.Player {
	return d.Players[index]
}

func (d *Domain) GetDefaultPlayer() konstructor.Player {
	return d.Players[0]
}
