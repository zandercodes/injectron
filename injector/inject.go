package injector

import (
	"encoding/json"
	"log"
	"time"

	"github.com/rgamba/evtwebsocket"
)

// Inject is the main function that injects the scripts into the executable
func (i *Injector) Inject() error {
	if !i.checkExecutable() {
		return ErrExecutableNotFound
	}

	if i.killIfRunning {
		if err := i.killProcess(); err != nil {
			return err
		}
	}

	if err := i.waitingOfApp(); err != nil {
		return err
	}

	time.Sleep(time.Duration(i.waitingBeforeInject) * time.Millisecond)

	sockets, err := i.getSockets()
	if err != nil {
		return err
	}

	for _, socket := range sockets {
		var script string
		for _, s := range i.scripts {
			script += s
		}
		if err := evalScript(socket, script); err != nil {
			return err
		}
	}

	return nil
}

func evalScript(socket string, script string) error {
	c := evtwebsocket.Conn{}
	err := c.Dial(socket, "")

	var request map[string]interface{} = make(map[string]interface{})
	request["id"] = 1
	request["method"] = "Runtime.evaluate"
	request["params"] = map[string]interface{}{
		"expression":            script,
		"objectGroup":           "inject",
		"includeCommandLineAPI": true,
		"silent":                true,
		"userGesture":           true,
		"awaitPromise":          true,
	}
	out, err := json.Marshal(request)
	if err != nil {
		return err
	}

	msg := evtwebsocket.Msg{
		Body: out,
		Callback: func(b []byte, c *evtwebsocket.Conn) {
			log.Println(string(b))
		},
	}

	c.Send(msg)
	return nil
}
