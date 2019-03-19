package main

// Collider @todo should this take reference to the collider that its touching
type Collider struct {
	x         float64
	y         float64
	w         float64
	h         float64
	colliding bool
}

// Intersects checks the x and y value for intersection to passed collider
func (c *Collider) Intersects(ac *Collider) bool {
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

func (c *Collider) Translate(x float64, y float64) {
	c.x += x
	c.y += y
}
