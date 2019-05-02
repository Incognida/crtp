//go:generate genny -in=/home/zaur/crtp/template/rectangle.go -out=/home/zaur/crtp/gen/generated.go gen "Item=Frame"
package main

import (
	"crtp/gen"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

func prettyPrint(info string, obj interface{}) {
	fmt.Println(strings.ToUpper(info))
	spew.Dump(obj)
	fmt.Println()
}

func main() {
	r := &gen.Rectangle{}
	f := &gen.Frame{}
	// embed this to get ability to polymorph SetWidth, SetHeight
	f.Rectangle = r
	// this one needed to return instance of frame, instead of casting to generic type e.g. return (T)this
	r.Generic = f

	prettyPrint("current frame", f)

	s1 := f.SetWidth(10)
	prettyPrint("after setting width", s1)

	s2 := s1.SetHeight(5)
	prettyPrint("after setting height", s2)

	s3 := s2.SetColor("blue")
	prettyPrint("after setting color", s3)
}