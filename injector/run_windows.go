//go:build windows

package injector

import (
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

// startProcess starts the process
func (i *Injector) startProcess() error {
	h, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, uint32(os.Getpid()))
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(h)

	var token syscall.Token
	err = syscall.OpenProcessToken(h, syscall.TOKEN_ALL_ACCESS, &token)
	if err != nil {
		return err
	}

	cmd := exec.Command(i.executable)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Token: token,
	}
	cmd.Args = append(cmd.Args, "--remote-debugging-port="+strconv.Itoa(i.debugPort))

	err = cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
