package grifts

import (
	"github.com/admin/sentryo/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
