package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Entity struct containing logic for game entites
type Entity struct {
	Image    *ebiten.Image
	GeoM     *ebiten.GeoM
	Collider *Collider
	DrawOpts *ebiten.DrawImageOptions
}

// NewEntity configures an entity at an x/y position
func NewEntity(x float64, y float64) *Entity {
	entityGeom := ebiten.GeoM{}
	entityGeom.Translate(x, y)

	drawOpts := ebiten.DrawImageOptions{GeoM: entityGeom}

	e, _ := ebiten.NewImage(100, 100, ebiten.FilterDefault)
	e.Fill(color.White)

	return &Entity{e, &drawOpts.GeoM, &Collider{x: x, y: y, w: 100, h: 100, col: &Collision{false, false, false, false}}, &drawOpts}
}

func (e *Entity) Translate(x float64, y float64) {
	e.GeoM.Translate(x, y)
	e.Collider.Translate(x, y)
}

func (e *Entity) SetColor(col color.RGBA) {
	e.Image.Fill(col)
}
