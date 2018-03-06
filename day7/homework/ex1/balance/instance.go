package balance

import (
	"strconv"
)

type Instance struct {
	host string
	port int
}

func InitInstance(host string, port int) *Instance {

	return &Instance{
		host: host,
		port: port,
	}
}

func (this *Instance) GetHost() string {
	return this.host
}

func (this *Instance) GetPort() int {
	return this.port
}

func (this *Instance) String() string {
	return this.GetHost() + ":" + strconv.Itoa(this.GetPort())
}
