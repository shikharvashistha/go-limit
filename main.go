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
	return nil
}

func main() {
	all := All{}
	all.limitAll()
}
