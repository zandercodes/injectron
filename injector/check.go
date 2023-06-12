package injector

// checkExecutable checks if the executable passed in is valid
func (i *Injector) checkExecutable() bool {
	if fileInfo, err := getFileInfo(i.executable); err == nil {
		return !fileInfo.IsDir()
	}
	return false
}
