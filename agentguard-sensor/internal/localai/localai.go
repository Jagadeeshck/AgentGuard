package localai

import (
	"fmt"
	"net"
)

type Service struct {
	Port    int
	Addr    string
	Exposed bool
}

var Ports = []int{11434, 1234, 7860, 8000, 8080, 3000, 5000, 5173, 8888}

func Scan() []Service {
	out := []Service{}
	for _, p := range Ports {
		for _, a := range []string{"127.0.0.1", "0.0.0.0"} {
			c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", a, p))
			if err == nil {
				_ = c.Close()
				out = append(out, Service{Port: p, Addr: a, Exposed: a != "127.0.0.1"})
			}
		}
	}
	return out
}
