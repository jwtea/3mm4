package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game struct
type Game struct {
	p      Player
	g      *Graphics
	img    *ebiten.Image
	target *Entity
}

//NewGame returns a new game struct
func NewGame() Game {
	return Game{p: NewPlayer(0, 0), g: &Graphics{800, 600}, target: NewEntity(300.0, 300.0)}
}

//Start begins game logic
func (g *Game) Start() {
	log.Printf("Starting Game")
	ebiten.Run(g.update, g.g.Height, g.g.Width, 2, "3mm4")
}

func (g *Game) update(screen *ebiten.Image) error {
	g.p.Controller.HandleInputs()

	// Background
	img, _ := ebiten.NewImage(g.g.Height, g.g.Width, ebiten.FilterDefault)

	img.Fill(color.Black)
	screen.DrawImage(img, &ebiten.DrawImageOptions{})

	screen.DrawImage(g.p.Image, g.p.DrawOpts)
	screen.DrawImage(g.target.Image, g.target.DrawOpts)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("3mm4: %0.2f", ebiten.CurrentTPS()))

	g.checkColliders()
	return nil
}

func (g *Game) checkColliders() {
	collided := color.RGBA{
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(0xff),
	}
	if g.p.Collider.Intersects(g.target.Collider) {
		g.p.SetColor(collided)
	}
}
