package internal

import (
	"github.com/kevincoding525/leaf/gate"
	"github.com/kevincoding525/leaf/log"
	"reflect"
	"server/msg"
)

func init() {
	handler(&msg.Hello{}, handlerHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)
	log.Debug("hello %v", m.Name)
	a.WriteMsg(&msg.Hello{Name: "client"})
}
