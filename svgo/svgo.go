package svgo

import (
	"github.com/skelterjohn/go.vg"
	"github.com/ajstarks/svgo"
	"math"
	"fmt"
	"io"
)

func Render(w io.Writer, proc *vg.Procedure, width, height int) {
	svg := svg.New(w)

	svg.Start(width, height)
	render(svg, proc.Operations)
	svg.End()
}

func render(svg *svg.SVG, ops []vg.Operation) {
	var line, fill vg.Color
	colorStyle := "fill:black; stroke:black"
	updateColorStyle := func() {
		fa := float64(fill.A)/255
		la := float64(line.A)/255
		colorStyle = fmt.Sprintf("fill:rgb(%d,%d,%d); fill-opacity:%f; stroke:rgb(%d,%d,%d); stroke-opacity:%f",
			fill.R, fill.G, fill.B, fa, line.R, line.G, line.B, la)
	}
	for _, op := range ops {
		switch op.Kind {
		case vg.OpTranslate:
			x, y := int(op.Floats[0]), int(op.Floats[1])
			svg.Translate(x, y)
		case vg.OpRotate:
			deg := op.Floats[0] * 180 / math.Pi
			svg.Rotate(deg)
		case vg.OpScale:
			sx, sy := op.Floats[0], op.Floats[1]
			svg.ScaleXY(sx, sy)
		case vg.OpLineColor:
			r, g, b, a := op.Ints[0], op.Ints[1], op.Ints[2], op.Ints[3]
			line = vg.Color{r, g, b, a}
			updateColorStyle()
		case vg.OpFillColor:
			r, g, b, a := op.Ints[0], op.Ints[1], op.Ints[2], op.Ints[3]
			fill = vg.Color{r, g, b, a}
			updateColorStyle()
		case vg.OpLine:
			x1, y1, x2, y2 := int(op.Floats[0]), int(op.Floats[1]), int(op.Floats[2]), int(op.Floats[3])
			svg.Line(x1, y1, x2, y2, colorStyle)
		case vg.OpRect:
			x, y, w, h := int(op.Floats[0]), int(op.Floats[1]), int(op.Floats[2]), int(op.Floats[3])
			left := x - w/2
			upper := y - h/2
			svg.Rect(left, upper, w, h, colorStyle)
		case vg.OpEllipse:
			x, y, w, h := int(op.Floats[0]), int(op.Floats[1]), int(op.Floats[2]), int(op.Floats[3])
			svg.Ellipse(x, y, w/2, h/2, colorStyle)
		case vg.OpPolygon:
			xs := make([]int, len(op.Floats)/2)
			ys := make([]int, len(op.Floats)/2)
			for i:=0; i<len(xs); i++ {
				xs[i] = int(op.Floats[i*2])
				ys[i] = int(op.Floats[i*2+1])
			}
			svg.Polygon(xs, ys, colorStyle)
		case vg.OpProc:
			render(svg, op.Operations)
		}
	}
}