//go:build !windows

package injector

func (i *Injector) Inject() error {
	// TODO: Implement this for linux
	return errors.New("not implemented yet for linux")
}
