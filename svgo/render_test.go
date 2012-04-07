package svgo

import (
	"testing"
	"github.com/skelterjohn/go.vg"
	"os"
)

func TestSVGo(t *testing.T) {
	proc := &vg.Procedure{}
	proc.FillColor(vg.Color{0, 127, 0, 255})
	proc.LineColor(vg.Color{255, 0, 0, 255})
	proc.Rect(vg.Point{50, 50}, vg.Size{50, 50})
	proc.LineColor(vg.Color{0, 0, 255, 255})
	proc.FillColor(vg.Color{255, 255, 255, 127})
	proc.Ellipse(vg.Point{50, 50}, vg.Size{50, 50})
	Render(os.Stdout, proc, 100, 100)
}
