package vg

type Color struct {
	A, R, G, B uint8
}

type Point struct {
	X, Y float64
}

type Size struct {
	W, H float64
}

type Procedure struct {
	operations []operation
}

func (me *Procedure) Translate(offset Point) {
	me.operations = append(me.operations, operation{
		kind: opTranslate,
		floats: []float64{offset.X, offset.Y},
	})
}

func (me *Procedure) Rotate(radians float64) {
	me.operations = append(me.operations, operation{
		kind: opRotate,
		floats: []float64{radians},
	})
}

func (me *Procedure) Scale(size Size) {
	me.operations = append(me.operations, operation{
		kind: opScale,
		floats: []float64{size.W, size.H},
	})
}

func (me *Procedure) LineColor(color Color) {
	me.operations = append(me.operations, operation{
		kind: opLineColor,
		ints: []uint8{color.A, color.R, color.G, color.B},
	})
}

func (me *Procedure) FillColor(color Color) {
	me.operations = append(me.operations, operation{
		kind: opFillColor,
		ints: []uint8{color.A, color.R, color.G, color.B},
	})
}

func (me *Procedure) Line(start, end Point) {
	me.operations = append(me.operations, operation{
		kind: opLine,
		floats: []float64{start.X, start.Y, end.X, end.Y},
	})
}

func (me *Procedure) Rect(center Point, size Size) {
	me.operations = append(me.operations, operation{
		kind: opRect,
		floats: []float64{center.X, center.Y, size.W, size.H,},
	})
}

func (me *Procedure) Ellipse(center Point, size Size) {
	me.operations = append(me.operations, operation{
		kind: opEllipse,
		floats: []float64{center.X, center.Y, size.W, size.H,},
	})
}

func (me *Procedure) Poly(vertices []Point) {
	op := operation{
		kind: opPolygon,
	}
	me.operations = append(me.operations, op)
	for _, v := range vertices {
		op.floats = append(op.floats, v.X, v.Y)
	}
}

func (me *Procedure) Proc(proc *Procedure) {
	me.operations = append(me.operations, operation{
		kind: opProc,
		operations: proc.operations,
	})
}

func (me *Procedure) ProcClip(proc *Procedure, center Point, size Size) {
	me.operations = append(me.operations, operation{
		kind: opProc,
		floats: []float64{center.X, center.Y, size.W, size.H,},
		operations: proc.operations,
	})
}
