package filters

import (
	"image/color"
	"testing"

	"github.com/DavidEsdrs/image-processing/quad"
)

func TestTurnLeft(t *testing.T) {
	t.Run("when turning left, the height and width must be swapped", func(t *testing.T) {
		q := quad.NewQuad(3, 1)

		if err := q.SetPixel(0, 0, color.RGBA{0, 127, 255, 255}); err != nil {
			t.Fatalf("failed to set pixel at (0, 0): %v", err)
		}

		if err := q.SetPixel(1, 0, color.RGBA{0, 127, 255, 255}); err != nil {
			t.Fatalf("failed to set pixel at (1, 0): %v", err)
		}

		if err := q.SetPixel(2, 0, color.RGBA{0, 127, 255, 255}); err != nil {
			t.Fatalf("failed to set pixel at (2, 0): %v", err)
		}

		tl := NewTurnLeftFilter()

		if err := tl.Execute(q); err != nil {
			t.Fatalf("failed to execute TurnLeftFilter: %v", err)
		}

		if q.Rows != 3 {
			t.Fatalf("unexpected height after turning left: got %d, want %d", q.Rows, 3)
		}

		if q.Cols != 1 {
			t.Fatalf("unexpected width after turning left: got %d, want %d", q.Cols, 1)
		}

		expectedColor := color.RGBA{0, 127, 255, 255}

		if got := q.GetPixel(0, 0); got != expectedColor {
			t.Fatalf("pixel at (0, 0) has incorrect color: got %v, want %v", got, expectedColor)
		}

		if got := q.GetPixel(0, 1); got != expectedColor {
			t.Fatalf("pixel at (0, 1) has incorrect color: got %v, want %v", got, expectedColor)
		}

		if got := q.GetPixel(0, 2); got != expectedColor {
			t.Fatalf("pixel at (0, 2) has incorrect color: got %v, want %v", got, expectedColor)
		}
	})
}
