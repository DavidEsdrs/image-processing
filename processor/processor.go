package processor

import (
	"image/color"

	"github.com/DavidEsdrs/image-processing/quad"
)

type Command interface {
	Execute(*quad.Quad) error
}

type Invoker struct {
	processes      []Command
	ColorModel     color.Model
	FiltersApplied int
}

// Invoke invokes all the commands (processes, filters or transformations) and
// applies it to the given tensor, which represents the input image.
func (i *Invoker) Invoke(q *quad.Quad) error {
	for _, p := range i.processes {
		err := p.Execute(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Invoker) GetColorModel() color.Model {
	return i.ColorModel
}

func (i *Invoker) SetColorModel(colorModel color.Model) {
	i.ColorModel = colorModel
}

func (i *Invoker) AddProcess(c Command) {
	i.processes = append(i.processes, c)
	i.FiltersApplied++
}

func (i *Invoker) ShouldInvoke() bool {
	return i.FiltersApplied > 0
}
