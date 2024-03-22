package command

import "github.com/smiling77877/coredemo/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(initCronCommand())
	root.AddCommand(initAppCommand())
}
