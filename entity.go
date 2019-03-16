package main

import "github.com/hajimehoshi/ebiten"

// Entity struct containing logic for game entites
type Entity struct {
	*ebiten.Image
	*ebiten.GeoM
	*Collider
	*ebiten.DrawImageOptions
}
