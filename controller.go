package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

//Controller contains a target GeoM and movements to apply
type Controller struct {
	Target *ebiten.GeoM
}

// HandleUpPress returns wether the user has pressed down up
func HandleUpPress() {
	log.Println("Key up pressed")
}

//HandleInputs check for keypresses and apply translations
func (c *Controller) HandleInputs() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.Up()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.Down()
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.Left()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.Right()
	}
}

// Up apply translate to geom target x-,y-10
func (c *Controller) Up() {
	log.Println("Up")
	c.Target.Translate(0, -10)
}

// Down apply translate to geom target x, y10
func (c *Controller) Down() {
	c.Target.Translate(0, 10)
}

// Left apply transltate x-10, y
func (c *Controller) Left() {
	c.Target.Translate(-10, 0)
}

// Right apply translate x10,y
func (c *Controller) Right() {
	c.Target.Translate(10, 0)
}
