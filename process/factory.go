package process

type Factory interface {
	NewProcess(name string, args ...string) *Process
}

type NativeProcessFactory struct {
}

func NewFactory() *NativeProcessFactory {
	return &NativeProcessFactory{}
}

func (f *NativeProcessFactory) NewProcess(name string, args ...string) *Process {
	return New(name, args...)
}
