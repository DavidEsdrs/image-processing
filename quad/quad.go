package quad

import (
	"errors"
	"image/color"
)

// Stores pixels in a linear representation while offers a 2d interface
type Quad struct {
	pix        []uint8
	Cols, Rows int
}

func NewQuad(Cols, Rows int) *Quad {
	return &Quad{
		pix:  make([]uint8, Rows*Cols*4),
		Rows: Rows,
		Cols: Cols,
	}
}

func (q *Quad) Get(x, y int) (uint8, error) {
	if x >= q.Cols || y >= q.Rows {
		return 0, errors.New("invalid access")
	}
	idx := (y * q.Cols) + x
	return q.pix[idx], nil
}

func (q *Quad) GetSlice(x, y, size int) ([]uint8, error) {
	if x >= q.Cols*4 || y >= q.Rows {
		return []uint8{}, errors.New("invalid access")
	}
	idx := (y * q.Cols) + x
	return q.pix[idx : idx+size : idx+size], nil
}

func (q *Quad) Set(x, y int, b uint8) error {
	if x >= q.Cols || y >= q.Rows {
		return errors.New("invalid access")
	}
	idx := (y * q.Cols) + x
	q.pix[idx] = b
	return nil
}

func (q *Quad) SetPixel(x, y int, rgb color.RGBA) {
	s, err := q.GetSlice(x*4, y, 4)
	if err != nil {
		panic("unexpected error while getting slice")
	}
	s[0] = rgb.R
	s[1] = rgb.G
	s[2] = rgb.B
	s[3] = rgb.A
}

func (q *Quad) GetPixel(x, y int) color.RGBA {
	s, err := q.GetSlice(x*4, y, 4)
	if err != nil {
		panic("unexpected error while getting slice")
	}
	return color.RGBA{s[0], s[1], s[2], s[3]}
}

func (q *Quad) Iterate(f func(color.RGBA)) {
	for idx := 0; idx < len(q.pix); idx += 4 {
		channels := q.pix[idx : idx+3 : idx+3]
		pixel := color.RGBA{channels[0], channels[1], channels[2], channels[3]}
		f(pixel)
	}
}

// iterate over each pixel applying the given filter onto it
func (q *Quad) Apply(filter func(color.RGBA) color.RGBA) {
	for idx := 0; idx < len(q.pix); idx += 4 {
		channels := q.pix[idx : idx+3 : idx+3]
		pixel := color.RGBA{channels[0], channels[1], channels[2], channels[3]}
		newPixel := filter(pixel)
		channels[0] = newPixel.R
		channels[1] = newPixel.G
		channels[2] = newPixel.B
		channels[3] = newPixel.A
	}
}

// returns a new *Quad with the same pixels
func (q *Quad) Clone() *Quad {
	quadClone := NewQuad(q.Cols, q.Rows)
	copy(quadClone.pix, q.pix)
	return quadClone
}
