package vg

type Color struct {
	R, G, B, A uint8
}

type Point struct {
	X, Y float64
}

type Size struct {
	W, H float64
}

type Procedure struct {
	Operations []Operation
}

func (me *Procedure) Translate(offset Point) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpTranslate,
		Floats: []float64{offset.X, offset.Y},
	})
}

func (me *Procedure) Rotate(radians float64) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpRotate,
		Floats: []float64{radians},
	})
}

func (me *Procedure) Scale(size Size) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpScale,
		Floats: []float64{size.W, size.H},
	})
}

func (me *Procedure) LineColor(color Color) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpLineColor,
		Ints: []uint8{color.R, color.G, color.B, color.A},
	})
}

func (me *Procedure) FillColor(color Color) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpFillColor,
		Ints: []uint8{color.R, color.G, color.B, color.A},
	})
}

func (me *Procedure) Line(start, end Point) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpLine,
		Floats: []float64{start.X, start.Y, end.X, end.Y},
	})
}

func (me *Procedure) Rect(center Point, size Size) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpRect,
		Floats: []float64{center.X, center.Y, size.W, size.H,},
	})
}

func (me *Procedure) Ellipse(center Point, size Size) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpEllipse,
		Floats: []float64{center.X, center.Y, size.W, size.H,},
	})
}

func (me *Procedure) Poly(vertices []Point) {
	op := Operation{
		Kind: OpPolygon,
	}
	me.Operations = append(me.Operations, op)
	for _, v := range vertices {
		op.Floats = append(op.Floats, v.X, v.Y)
	}
}

func (me *Procedure) Proc(proc *Procedure) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpProc,
		Operations: proc.Operations,
	})
}

func (me *Procedure) ProcClip(proc *Procedure, center Point, size Size) {
	me.Operations = append(me.Operations, Operation{
		Kind: OpProc,
		Floats: []float64{center.X, center.Y, size.W, size.H,},
		Operations: proc.Operations,
	})
}
