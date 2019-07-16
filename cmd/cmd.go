package cmd

import (
	"go.zoe.im/x/cli"
)

var (
	// global cmd contains all sub command
	cmd = cli.New(
		cli.Name("goser"),
		cli.Short("Goser is a toolbox to generate for common backend develop."),
		cli.Description(`Generate code from yaml to RPC, REST API, database etc.
		
Default we will generate code in current directory,
and use parent directory name as package name.

To make the code directory clear you can put all the defined proto yaml
files to a directory.

But if you pointer a direcotiry for output, we will use the output
directory name as the package name.

Certainly, you can use a flag --pkg to set your package name.
		`),
		cli.Run(func(c *cli.Command, args ...string) {
			// we at here to start create project
			generate(args...)
		}),
	)
)

// Register create a sub command
func Register(cmds ...*cli.Command) error {
	return cmd.Register(cmds...)
}

// Run execute command
func Run(opts ...cli.Option) error {
	return cmd.Run(opts...)
}
