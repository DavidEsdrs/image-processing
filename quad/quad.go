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

// return the underlying slice
func (q *Quad) GetUnderlyingSlice() []uint8 {
	return q.pix
}

// Método que corta um número especificado de linhas do topo da imagem
func (q *Quad) CropTop(lines int) error {
	if lines >= q.Rows {
		return errors.New("cannot crop more rows than present")
	}

	// Atualiza a quantidade de linhas e fatia o array de pixels para remover as linhas do topo
	q.pix = q.pix[lines*q.Cols*4:] // Avança o ponteiro de pixels ignorando as primeiras `lines` linhas
	q.Rows -= lines                // Ajusta o número total de linhas
	return nil
}

// Método que corta um número especificado de colunas do lado esquerdo da imagem
func (q *Quad) CropLeft(columns int) error {
	if columns >= q.Cols {
		return errors.New("cannot crop more columns than present")
	}

	// Ajusta o array de pixels para remover as colunas do lado esquerdo
	for y := 0; y < q.Rows; y++ {
		start := (y*q.Cols + columns) * 4 // Ponto inicial para cada linha após remover colunas
		end := (y + 1) * q.Cols * 4       // Fim da linha original
		q.pix = append(q.pix[:start], q.pix[start:end]...)
	}

	q.Cols -= columns // Ajusta o número de colunas
	return nil
}

func (q *Quad) Get(x, y int) (uint8, error) {
	if x >= q.Cols || y >= q.Rows {
		return 0, errors.New("invalid access")
	}
	idx := (y * q.Cols) + x
	return q.pix[idx], nil
}

func (q *Quad) GetSlice(x, y, size int) ([]uint8, error) {
	// Verificar se x e y estão dentro dos limites apropriados de coluna e linha
	if x < 0 || y < 0 || x+size > q.Cols*4 || y >= q.Rows {
		return []uint8{}, errors.New("invalid access")
	}

	// Corrigir o cálculo do índice para incluir a largura de 4 canais por pixel
	idx := (y * q.Cols * 4) + x

	// Verificar se o slice resultante está dentro dos limites de q.pix
	if idx+size > len(q.pix) {
		return []uint8{}, errors.New("invalid slice range")
	}

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

func (q *Quad) SetPixel(col, row int, rgb color.RGBA) error {
	if col < 0 || col >= q.Cols*4 || row < 0 || row >= q.Rows {
		return errors.New("coordinates out of bounds")
	}

	idx := (row * q.Cols * 4) + (col * 4)

	q.pix[idx] = rgb.R
	q.pix[idx+1] = rgb.G
	q.pix[idx+2] = rgb.B
	q.pix[idx+3] = rgb.A
	return nil
}

func (q *Quad) GetPixel(x, y int) color.RGBA {
	s, err := q.GetSlice(x*4, y, 4)
	if err != nil {
		panic("unexpected error while getting slice")
	}
	return color.RGBA{s[0], s[1], s[2], s[3]}
}

// GetRow obtém uma linha do Quad como um slice de color.Color.
func (q *Quad) GetRow(y int) []color.RGBA {
	row := make([]color.RGBA, q.Cols)
	for x := 0; x < q.Cols; x++ {
		r, g, b, a := q.GetPixel(x, y).R, q.GetPixel(x, y).G, q.GetPixel(x, y).B, q.GetPixel(x, y).A
		row[x] = color.RGBA{r, g, b, a}
	}
	return row
}

// SetRow define uma linha no Quad com um slice de color.Color.
func (q *Quad) SetRow(y int, row []color.RGBA) {
	for x := 0; x < q.Cols; x++ {
		if x < len(row) {
			q.SetPixel(x, y, row[x])
		}
	}
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
		channels := q.pix[idx : idx+4 : idx+4]
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
