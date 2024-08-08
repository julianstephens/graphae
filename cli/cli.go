package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type VersionFlag string

type Globals struct {
	Debug    bool        `short:"D" help:"Enable debug mode"`
	LogLevel string      `short:"l" help:"Set the logging level (debug|info|warn|error|fatal)" default:"info"`
	Version  VersionFlag `name:"version" help:"Print version information and quit"`
}

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type CLI struct {
	Globals

	Transform TransformCmd `cmd:"" help:"Transform an image"`
}
