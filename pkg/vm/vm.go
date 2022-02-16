package vm

import "fmt"

type vmApi struct {
	ab func()
	cd func()
}

func (v *vmApi) Limit() error {
	v.ab() = func() {
		fmt.Println("vm.Limit()")
	}
	v.cd() = func() {
		fmt.Println("vm.Limit()")
	}
	return nil
}

func main() {
	g := &vmApi{}
	g.Limit()
}
