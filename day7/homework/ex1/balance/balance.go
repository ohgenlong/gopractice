package balance

type BaLance interface {
	DoBalance(i []*Instance) (*Instance, error)
}
