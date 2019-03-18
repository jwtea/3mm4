package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Player struct containing all data pertaining to the player
type Player struct {
	*Entity
	*Controller
}

// NewPlayer returns a configured player object
func NewPlayer(x float64, y float64) Player {
	// Create a geom for the player and initalise location
	playerGeom := ebiten.GeoM{}
	playerGeom.Translate(x, y)

	drawOpts := ebiten.DrawImageOptions{GeoM: playerGeom}

	p, _ := ebiten.NewImage(100, 100, ebiten.FilterDefault)
	p.Fill(color.White)

	controller := Controller{&drawOpts.GeoM}
	return Player{&Entity{p, &drawOpts.GeoM, &Collider{}, &drawOpts}, &controller}
}
