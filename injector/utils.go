package injector

import (
	"os"
)

func getFileInfo(executable string) (os.FileInfo, error) {
	return os.Stat(executable)
}
