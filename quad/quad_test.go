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

func TestSetPixel(t *testing.T) {
	t.Run("multidimensional quad must be stored successfully", func(t *testing.T) {
		q := NewQuad(3, 3)

		expected := [][]color.RGBA{
			{color.RGBA{100, 127, 255, 100}, color.RGBA{200, 100, 20, 0}, color.RGBA{127, 42, 37, 20}},
			{color.RGBA{0, 8, 128, 100}, color.RGBA{192, 0, 32, 57}, color.RGBA{127, 42, 37, 20}},
			{color.RGBA{255, 0, 0, 100}, color.RGBA{100, 40, 20, 40}, color.RGBA{99, 42, 11, 255}},
		}

		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				err := q.SetPixel(col, row, expected[row][col])
				if err != nil {
					t.Fatalf(err.Error())
				}
			}
		}

		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				pixel := q.GetPixel(col, row)
				if pixel != expected[row][col] {
					t.Fatalf("unexpected pixel got - expected: %v | actual: %v\n", expected[row][col], pixel)
				}
			}
		}
	})

}
