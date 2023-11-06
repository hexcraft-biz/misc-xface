package test

import (
	"fmt"
	"testing"

	xface "github.com/hexcraft-biz/misc-xface"
)

func TestXface(t *testing.T) {
	f1 := new(xface.Descriptor)
	f2 := new(xface.Descriptor)
	fmt.Println("Dist:", f1.DistWithFace(f2))
}
