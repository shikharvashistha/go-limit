package main

type GcpApis struct {
	vm.vmRateLimiter
	disk.diskRateLimiter
}

func (g *GcpApis) diskCost() {
}
