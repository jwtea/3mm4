package main

import "log"

// Collider @todo should this take reference to the collider that its touching
type Collider struct {
	x         float64
	y         float64
	w         float64
	h         float64
	colliding bool
	col       *Collision
}

// Collision object used to determine where a collision is occuring
type Collision struct {
	top    bool
	right  bool
	bottom bool
	left   bool
}

// Intersects checks the x and y value for intersection to passed collider
func (c *Collider) Intersects(ac *Collider) bool {

	c.collideBottom(ac)
	c.collideTop(ac)
	c.collideLeft(ac)
	c.collideRight(ac)

	if c.IntersectsX(ac) && c.IntersectsY(ac) {
		return true
	}
	return false
}

// IntersectsX checks for the width intersecting the x value of the passed collider
func (c *Collider) IntersectsX(ac *Collider) bool {
	if c.x+c.w >= ac.x && c.x <= ac.x+ac.w {
		return true
	}

	return false
}

// IntersectsY checks for the width intersecting the y value of the passed collider
func (c *Collider) IntersectsY(ac *Collider) bool {
	if c.y+c.h >= ac.y && c.y <= ac.y+ac.h {
		return true
	}

	return false
}

func (c *Collider) collideRight(ac *Collider) {
	if (c.x+c.w >= ac.x && c.x+c.w <= ac.x+ac.w) && c.IntersectsY(ac) {
		c.col.SetRight(true)
	} else {
		c.col.SetRight(false)
	}
}

func (c *Collider) collideLeft(ac *Collider) {
	if (c.x >= ac.x && c.x <= ac.x+ac.w) && c.IntersectsY(ac) {
		c.col.SetLeft(true)
	} else {
		c.col.SetLeft(false)
	}

}

func (c *Collider) collideTop(ac *Collider) {
	if (c.y <= ac.y+ac.h && c.y >= ac.y) && c.IntersectsX(ac) {
		c.col.SetTop(true)
	} else {
		c.col.SetTop(false)
	}

}

func (c *Collider) collideBottom(ac *Collider) {
	if (c.y+c.h >= ac.y && c.y+c.h <= ac.y+ac.h) && c.IntersectsX(ac) {
		c.col.SetBottom(true)
	} else {
		c.col.SetBottom(false)
	}

}

func (c *Collider) Translate(x float64, y float64) {
	c.x += x
	c.y += y
}

func (c *Collision) SetTop(b bool) {
	c.top = b
}

func (c *Collision) SetRight(b bool) {
	c.right = b
}

func (c *Collision) SetBottom(b bool) {
	c.bottom = b
}

func (c *Collision) SetLeft(b bool) {
	c.left = b
}

func (c *Collision) PrintDebug() {
	log.Printf("T: %v ", c.top)
	log.Printf("R: %v", c.right)
	log.Printf("B: %v", c.bottom)
	log.Printf("L: %v", c.left)
}
