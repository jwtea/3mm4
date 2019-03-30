package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Entity struct containing logic for game entites
type Entity struct {
	Image    *ebiten.Image
	GeoM     *ebiten.GeoM
	Collider *Collider
	DrawOpts *ebiten.DrawImageOptions
	History  *History
}

// NewEntity configures an entity at an x/y position
func NewEntity(x float64, y float64) *Entity {
	entityGeom := ebiten.GeoM{}
	entityGeom.Translate(x, y)

	drawOpts := ebiten.DrawImageOptions{GeoM: entityGeom}

	e, _ := ebiten.NewImage(50, 50, ebiten.FilterDefault)
	e.Fill(color.White)

	return &Entity{e, &drawOpts.GeoM, &Collider{x: x, y: y, w: 50, h: 50, col: &Collision{false, false, false, false}}, &drawOpts, &History{}}
}

func (e *Entity) Translate(x float64, y float64) {

	//@todo should history be saved here
	e.History.Movements = append(e.History.Movements, &Movement{x, y, time.Now()})

	e.GeoM.Translate(x, y)
	e.Collider.Translate(x, y)
}

func (e *Entity) SetColor(col color.RGBA) {
	e.Image.Fill(col)
}
