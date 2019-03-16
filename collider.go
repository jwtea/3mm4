package main

// Collider
type Collider struct {
	x float64
	y float64
	w float64
	h float64
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
	if c.x+c.w >= ac.x {
		return true
	}

	return false
}

// IntersectsY checks for the width intersecting the y value of the passed collider
func (c *Collider) IntersectsY(ac *Collider) bool {
	if c.y+c.h >= ac.y {
		return true
	}

	return false
}
