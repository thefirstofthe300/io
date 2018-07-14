package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/hegemone/io/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
