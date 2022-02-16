package main

type GcpApis struct {
	vm.rateLimiter
	disk.rateLimiter
}
