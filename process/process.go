package process

import (
	"os"
	"os/exec"
	"strings"
)

type Process struct {
	Name         string
	Arguments    []*Argument
	WSLArguments []string
}

type Argument struct {
	Value    string
	IsSecret bool
}

func New(name string, args ...string) *Process {
	arguments := make([]*Argument, len(args))

	for i, arg := range args {
		arguments[i] = &Argument{Value: arg}
	}

	return &Process{
		Name:      name,
		Arguments: arguments,
	}
}

func (p *Process) Run() error {
	cmd := p.NewCommand()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (p *Process) RunBash() error {
	cmd := p.NewBashCommand()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (p *Process) Output() (string, error) {
	result, err := p.NewCommand().Output()

	return string(result), err
}

func (p *Process) OutputBash() (string, error) {
	result, err := p.NewBashCommand().Output()

	return string(result), err
}

func (p *Process) NewCommand() *exec.Cmd {
	args := p.buildArguments(false)

	return exec.Command(args[0], args[1:]...)
}

func (p *Process) NewBashCommand() *exec.Cmd {
	args := p.buildBashArguments(false)

	return exec.Command(args[0], args[1:]...)
}

func (p *Process) SetSecretArguments(indexes ...int) {
	for _, index := range indexes {
		if 0 <= index && index < len(p.Arguments) {
			p.Arguments[index].IsSecret = true
		}
	}
}

func (p *Process) String() string {
	return strings.Join(p.buildArguments(true), " ")
}

func (p *Process) BashString() string {
	return strings.Join(p.buildBashArguments(true), " ")
}

func (p *Process) buildArguments(hideSecret bool) []string {
	args := make([]string, len(p.WSLArguments)+1+len(p.Arguments))
	index := 0

	for _, arg := range p.WSLArguments {
		args[index] = arg
		index++
	}

	args[index] = p.Name
	index++

	for _, arg := range p.Arguments {
		if true == arg.IsSecret && true == hideSecret {
			args[index] = "***secret***"
		} else {
			args[index] = arg.Value
		}
		index++
	}

	return args
}

func (p *Process) buildBashArguments(hideSecret bool) []string {
	args := make([]string, len(p.WSLArguments)+3)
	index := 0

	for _, arg := range p.WSLArguments {
		args[index] = arg
		index++
	}

	args[index] = "bash"
	args[index+1] = "-c"

	args2 := make([]string, 1+len(p.Arguments))
	args2[0] = p.Name
	index2 := 1

	for _, arg := range p.Arguments {
		if true == arg.IsSecret && true == hideSecret {
			args2[index2] = "***secret***"
		} else {
			args2[index2] = arg.Value
		}
		index2++
	}

	args[index+2] = strings.Join(args2, " ")

	return args
}
