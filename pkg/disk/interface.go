package disk

import "fmt"

type DiskApi interface {
	ab() func() //respective functions
	cd() func()
}
type Final struct {
	diskLimiter     func() error
	diskListLimiter func() error
}

func (d *Final) ab() func() {
	return func() {
		//limit this function according to the requirement
		fmt.Println("disk ab")
	}
}
func (d *Final) cd() func() {
	return func() {
		//limit this function
		fmt.Println("disk cd")
	}
}
