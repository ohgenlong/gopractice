package balance

import "fmt"

type BalanceMgr struct {
	allBalancer map[string]BaLance
}

var mgr BalanceMgr = BalanceMgr{
	allBalancer: make(map[string]BaLance),
}

func (this *BalanceMgr) register(name string, b BaLance) {
	this.allBalancer[name] = b
}

func RegisterBalance(name string, b BaLance) {
	mgr.register(name, b)
}

func DoBalance(name string, inSs []*Instance) (ins *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	fmt.Printf("1\n")
	if !ok {
		err = fmt.Errorf("Not found %s\n", name)
		return
	}
	ins, err = balancer.DoBalance(inSs)
	return

}
