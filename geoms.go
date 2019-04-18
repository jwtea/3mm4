package main

import "github.com/hajimehoshi/ebiten"

type Triangles struct {
	Verts *[]ebiten.Vertex
	Idx   *[]uint16
}

var tri = [][]int{
	{0, 0},
	{0, 1},
	{1, 0},
}

var triIdx = []uint16{0, 1, 2}
