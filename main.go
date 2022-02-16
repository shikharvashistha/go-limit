package main

import (
	"github.com/shikharvashistha/go-limit/pkg/disk"
	"github.com/shikharvashistha/go-limit/pkg/vm"
)

type All struct {
	disk.DiskApi
	vm.VmApi
}

func (d *All) limitAll() error {
	//limit the whole abstraction here
	return nil
}

func main() {
	all := All{}
	all.limitAll()
}
