package disk

type diskApi interface {
	listDisk() error
	createDisk() error
	deleteDisk() error
}