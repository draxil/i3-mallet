package i3

import (
	"fmt"

	"go.i3wm.org/i3/v4"
)

func FocusedWin() (string, error) {
	t, err := i3.GetTree()
	if err != nil {
		return "", err
	}

	fn := t.Root.FindFocused(func(n *i3.Node) bool {
		return n != nil && n.Focused
	})

	if fn == nil {
		return "", nil
	}

	id := fmt.Sprintf("%d", fn.ID)

	return id, nil
}
