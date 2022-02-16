package disk

import "fmt"

type DiskApi interface {
	ab() func()
	cd() func()
}
type Final struct {
	diskLimiter     func() error
	diskListLimiter func() error
}

func (d *Final) ab() func() {
	return func() {
		//limit this function
		fmt.Println("disk ab")
	}
}
func (d *Final) cd() func() {
	return func() {
		//limit this function
		fmt.Println("disk cd")
	}
}
