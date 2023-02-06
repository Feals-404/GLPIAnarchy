package cmd

import (
	"github.com/Feals-404/GLPIAnarchy/exploits"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:     "scan",
	Short:   "Scans the GLPI : Recovers version and potential vulnerability",
	Example: `GLPIAnarchy scan https://glpi.example.com`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := exploits.NormalizeURL(args[0])
		if exploits.CheckIsGLPI(target) {
			exploits.GetCVE(target)
			exploits.FindFilesFolder(target)
		}

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
