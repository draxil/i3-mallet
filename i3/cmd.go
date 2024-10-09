package i3

import (
	"fmt"

	"go.i3wm.org/i3/v4"
)

func Run(cmdfmt string, arg ...any) error {
	cmd := fmt.Sprintf(cmdfmt, arg...)
	_, err := i3.RunCommand(cmd)
	if err != nil {
		return err
	}
	return nil
}
