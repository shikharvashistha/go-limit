package disk

type diskApi struct {
	ab func()
	cd func()
}

func (d *diskApi) Limit() error {
	d.ab() = func() {
	}
	d.cd() = func() {
	}
	return nil
}
func main() {
	g := &diskApi{}
	g.Limit()
}
