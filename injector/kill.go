package injector

import (
	"os"

	"github.com/zandercodes/injectron/injector/ps"
)

// killProcess kills the process with the process ID passed in
func (i *Injector) killProcess() error {
	fileInfo, err := getFileInfo(i.executable)
	if err != nil {
		return err
	}

	processes, err := ps.Processes()
	if err != nil {
		return err
	}

	for _, process := range processes {
		if process.Executable() == fileInfo.Name() {
			proc, err := os.FindProcess(process.Pid())
			if err != nil {
				return err
			}
			err = proc.Kill()
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}
