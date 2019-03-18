package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Entity struct containing logic for game entites
type Entity struct {
	*ebiten.Image
	*ebiten.GeoM
	*Collider
	*ebiten.DrawImageOptions
}

// NewEntity configures an entity at an x/y position
func NewEntity(x float64, y float64) *Entity {
	entityGeom := ebiten.GeoM{}
	entityGeom.Translate(x, y)

	drawOpts := ebiten.DrawImageOptions{GeoM: entityGeom}

	e, _ := ebiten.NewImage(100, 100, ebiten.FilterDefault)
	e.Fill(color.White)

	return &Entity{e, &drawOpts.GeoM, &Collider{x: x, y: y, w: 100, h: 1000}, &drawOpts}
}

func (e *Entity) SetColor(col color.RGBA64) {
	e.Fill(col)
}
