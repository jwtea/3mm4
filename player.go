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
func NewPlayer() Player {
	// Create a geom for the player and initalise location
	playerGeom := ebiten.GeoM{}
	playerGeom.Apply(2.0, 3.0)

	drawOpts := ebiten.DrawImageOptions{GeoM: playerGeom}

	p, _ := ebiten.NewImage(100, 100, ebiten.FilterDefault)
	p.Fill(color.White)
	controller := Controller{&drawOpts.GeoM}
	return Player{&Entity{p, &drawOpts.GeoM, &Collider{}, &drawOpts}, &controller}
}
