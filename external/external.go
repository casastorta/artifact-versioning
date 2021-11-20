package external

import (
	"os/exec"
)

type ICommand interface {
	Execute(string, ...string) ([]byte, error)
}

type TCommand struct {
	Command ICommand
}

func (c *TCommand) Execute(command string, params ...string) ([]byte, error) {
	cx, err := exec.Command(command, params...).CombinedOutput()
	return cx, err
}
