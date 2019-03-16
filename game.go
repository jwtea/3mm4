package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game struct
type Game struct {
	p   Player
	g   *Graphics
	img *ebiten.Image
}

//NewGame returns a new game struct
func NewGame() Game {
	return Game{p: NewPlayer(), g: &Graphics{800, 600}}
}

//Start begins game logic
func (g *Game) Start() {
	log.Printf("Starting Game")

	//@too move to graphics
	ebiten.Run(g.update, g.g.Height, g.g.Width, 2, "3mm4")
}

func (g *Game) update(screen *ebiten.Image) error {
	g.p.Controller.HandleInputs()

	// Background
	img, _ := ebiten.NewImage(g.g.Height, g.g.Width, ebiten.FilterDefault)
	img.Fill(color.Black)
	screen.DrawImage(img, &ebiten.DrawImageOptions{})

	screen.DrawImage(g.p.Image, g.p.DrawImageOptions)

	title := "3mm4"
	ebitenutil.DebugPrint(screen, title)
	return nil
}
