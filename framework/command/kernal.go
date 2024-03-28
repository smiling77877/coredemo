package command

import "github.com/smiling77877/coredemo/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(initAppCommand())
	root.AddCommand(initCronCommand())
	root.AddCommand(initEnvCommand())
	root.AddCommand(initConfigCommand())
	//build命令
	root.AddCommand(initBuildCommand())
	//go build
	root.AddCommand(goCommand)
	//npm build
	root.AddCommand(npmCommand)
}
