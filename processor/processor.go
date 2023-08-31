package processor

import "image/color"

type Command interface {
	Execute(*[][]color.Color) error
}

type Invoker struct {
	processes  []Command
	ColorModel color.Model
}

func (i *Invoker) Invoke(tensor *[][]color.Color) *[][]color.Color {
	for _, p := range i.processes {
		p.Execute(tensor)
	}
	return tensor
}

func (i *Invoker) GetColorModel() color.Model {
	return i.ColorModel
}

func (i *Invoker) SetColorModel(colorModel color.Model) {
	i.ColorModel = colorModel
}

func (i *Invoker) AddProcess(c Command) {
	i.processes = append(i.processes, c)
}
