// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package gen

type Rectangle struct {
	Generic       *Frame
	width, height uint
}

func (r *Rectangle) SetWidth(width uint) *Frame {
	r.width = width
	return r.Generic
}

func (r *Rectangle) SetHeight(height uint) *Frame {
	r.height = height
	return r.Generic
}

func (r *Rectangle) GetWidth() uint {
	return r.width
}

func (r *Rectangle) GetHeight() uint {
	return r.height
}