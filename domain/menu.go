package domain

import (
	"os"
)

func (d *Domain) CloseGame() {
}

func (d *Domain) LoadtGame() {
}

func (d *Domain) SaveGame() {
}

func (d *Domain) StartGame() {
}

func (d *Domain) ExitGame() {
	os.Exit(1)
}
