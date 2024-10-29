package quad

import (
	"image/color"
	"testing"
)

func TestQuadClone(t *testing.T) {
	q := NewQuad(3, 1)

	q.SetPixel(0, 0, color.RGBA{255, 255, 255, 255})
	q.SetPixel(1, 0, color.RGBA{255, 255, 255, 255})
	q.SetPixel(2, 0, color.RGBA{255, 255, 255, 255})

	clone := q.Clone()

	clone.SetPixel(0, 0, color.RGBA{0, 0, 0, 0})

	if q.GetPixel(0, 0) == clone.GetPixel(0, 0) {
		t.FailNow()
	}
}
