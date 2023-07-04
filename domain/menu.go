package domain

import (
	"fmt"
	"os"
)

func (d *Domain) CloseGame() {
	fmt.Println("Close game")
}

func (d *Domain) LoadtGame() {
	fmt.Println("Load game")
}

func (d *Domain) SaveGame() {
	fmt.Println("Save game")
}

func (d *Domain) StartGame() {
	fmt.Println("Start game")
}

func (d *Domain) ExitGame() {
	fmt.Println("Exit game")
	os.Exit(1)
}
