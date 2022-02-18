package vm

type vmApi interface {
	listVM() error
	createVM() error
	deleteVM() error
}