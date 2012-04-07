package vg

const (
	OpTranslate = iota
	OpRotate
	OpScale
	OpLineColor
	OpFillColor
	OpLine
	OpRect
	OpEllipse
	OpPolygon
	OpProc
)

type Operation struct {
	Kind       int
	Floats     []float64
	Ints       []uint8
	Operations []Operation
}
