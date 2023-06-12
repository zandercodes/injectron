//go:build linux

package injector

import (
	"os/exec"
	"strconv"
)

// startProcess starts the process
func (i *Injector) startProcess() error {
	cmd := exec.Command(i.executable)
	cmd.Args = append(cmd.Args, "--remote-debugging-port="+strconv.Itoa(i.debugPort))

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
