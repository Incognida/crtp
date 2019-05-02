package template

import "github.com/cheekybits/genny/generic"

type Item generic.Type

type Rectangle struct {
	Generic       *Item
	width, height uint
}


func (r *Rectangle) SetWidth(width uint) *Item{
	r.width = width
	return r.Generic
}

func (r *Rectangle) SetHeight(height uint) *Item{
	r.height = height
	return r.Generic
}


func (r *Rectangle) GetWidth() uint {
	return r.width
}


func (r *Rectangle) GetHeight() uint {
	return r.height
}


