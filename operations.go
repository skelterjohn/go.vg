package vg

const (
	opTranslate = iota
	opRotate
	opScale
	opLineColor
	opFillColor
	opLine
	opRect
	opEllipse
	opPolygon
	opVertex
	opProc
)

type operation struct {
	kind       int
	floats     []float64
	ints       []uint8
	operations []operation
}
