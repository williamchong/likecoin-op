package workflow

import "github.com/spf13/cobra"

var WorkflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Workflow commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
