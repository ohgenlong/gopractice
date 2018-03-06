package balance

import (
	"errors"
	"fmt"
)

type RoundRobinBalance struct {
	curIndex int
}

// init函数--默认会执行的函数
func init() {
	RegisterBalance("roundrobin", &RoundRobinBalance{})
	fmt.Printf("rrb balance register!\n")
}

func (this *RoundRobinBalance) DoBalance(inSs []*Instance) (*Instance, error) {

	if len(inSs) == 0 {
		err := errors.New("no instance")
		return nil, err
	}

	lens := len(inSs)

	if this.curIndex >= lens {
		this.curIndex = 0
	}

	i := inSs[this.curIndex]
	this.curIndex = (this.curIndex + 1) % lens
	return i, nil
}
