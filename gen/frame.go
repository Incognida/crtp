package gen

type Frame struct {
	*Rectangle
	color string
}

func (f *Frame) SetColor(color string) *Frame{
	f.color = color
	return f
}