package balance

import (
	"errors"
	"math/rand"
	"fmt"
)


type RandomBanlance struct {

}

// init函数--默认会执行的函数
func init() {
	RegisterBalance("random", &RandomBanlance{})
	fmt.Printf("random balance register!\n")
}


func (this *RandomBanlance) DoBalance(inSs []*Instance) (*Instance, error) {

	if len(inSs) == 0 {
		err := errors.New("no instance")
		return nil, err
	}

	lens := len(inSs)
	index := rand.Intn(lens)
	i := inSs[index]
	return i, nil
}


