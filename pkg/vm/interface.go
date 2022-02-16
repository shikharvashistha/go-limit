package vm

import "fmt"

type VmApi interface {
	ab() func()
	cd() func()
}
type Final struct {
	vmLimiter     func() error
	vmListLimiter func() error
}

func (d *Final) ab() func() {
	return func() {
		fmt.Println("disk ab")
	}
}
func (d *Final) cd() func() {
	return func() {
		fmt.Println("disk cd")
	}
}