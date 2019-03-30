package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

//Controller contains a target GeoM and movements to apply
type Controller struct {
	Target *Entity
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
	c.Target.Collider.col.PrintDebug()
	c.Target.Translate(0, -10)
	if c.Target.Collider.col.top == true {
		c.Target.Translate(0, 11)
	}
}

// Down apply translate to geom target x, y10
func (c *Controller) Down() {
	c.Target.Translate(0, 10)
	if c.Target.Collider.col.bottom == true {
		c.Target.Translate(0, -11)
	}
}

// Left apply transltate x-10, y
func (c *Controller) Left() {
	c.Target.Translate(-10, 0)
	if c.Target.Collider.col.left == true {
		c.Target.Translate(11, 0)
	}
}

// Right apply translate x10,y
func (c *Controller) Right() {
	c.Target.Translate(10, 0)
	if c.Target.Collider.col.right == true {
		c.Target.Translate(-11, 0)
	}
}

//@todo not sure this should live here but allow for a history translation to be reversed
func (c *Controller) ReverseTranslation(x float64, y float64) {
	c.Target.Translate(x*-1, y*-1)
}
