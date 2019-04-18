package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func Conv2DIntToVertex(verts [][]int) *[]ebiten.Vertex {
	vx := []ebiten.Vertex{}

	// Why are these being used
	// sx, sy := -a.Width/2, -a.Height/2

	for _, vert := range verts {
		vx = append(vx, ebiten.Vertex{
			float32(vert[0]),
			float32(vert[1]),
			float32(vert[0]),
			float32(vert[1]),
			155,
			255,
			155,
			1,
		})
	}
	return &vx
}

//ScaleVertexes apply scale to all points in an ebiten.Vertex
func ScaleVertexes(scale float32, verts *[]ebiten.Vertex) *[]ebiten.Vertex {
	for i := 0; i < len(*verts); i++ {
		(*verts)[i].DstX = (*verts)[i].DstX * scale
		(*verts)[i].DstY = (*verts)[i].DstY * scale
		(*verts)[i].SrcX = (*verts)[i].SrcX * scale
		(*verts)[i].SrcY = (*verts)[i].SrcY * scale
	}
	return verts
}

func drawRect(canvas *ebiten.Image, img *ebiten.Image, x, y, width, height float32, address ebiten.Address, msg string) {
	sx, sy := -width/2, -height/2
	vs := []ebiten.Vertex{
		{
			DstX:   x,
			DstY:   y,
			SrcX:   sx,
			SrcY:   sy,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   x + width,
			DstY:   y,
			SrcX:   sx + width,
			SrcY:   sy,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   x,
			DstY:   y + height,
			SrcX:   sx,
			SrcY:   sy + height,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   x + width,
			DstY:   y + height,
			SrcX:   sx + width,
			SrcY:   sy + height,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
	}
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = address
	canvas.DrawTriangles(vs, []uint16{0, 1, 2, 1, 2, 3}, img, op)

	ebitenutil.DebugPrintAt(canvas, msg, int(x), int(y)-16)
}

//DrawGrid Show a grid on screen
func DrawGrid(canvas *ebiten.Image, width int, height int, sections int) {
	//Draw vertical
	for i := 0; i < sections; i++ {
		vO := (width / sections) * i
		hO := (height / sections) * i
		ebitenutil.DrawLine(
			canvas,
			float64(vO), 0,
			float64(vO), float64(height),
			color.White)

		ebitenutil.DrawLine(
			canvas,
			0, float64(hO),
			float64(width), float64(hO),
			color.White)
	}
}
