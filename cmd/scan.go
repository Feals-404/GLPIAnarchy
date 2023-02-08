package cmd

import (
	"github.com/Feals-404/GLPIAnarchy/exploits"
	"github.com/spf13/cobra"
)

var check bool
var flagver string
var scanCmd = &cobra.Command{
	Use:     "scan",
	Short:   "Scans the GLPI : Recovers version and potential vulnerability",
	Example: `GLPIAnarchy scan https://glpi.example.com`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := exploits.NormalizeURL(args[0])
		if exploits.CheckIsGLPI(target) {
			exploits.GetCVE(target, check, flagver)
			exploits.FindFilesFolder(target)
		}

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().BoolVarP(&check, "check", "c", false, "Perform active check to know if the target is vulnerable")
	scanCmd.Flags().StringVarP(&flagver, "glpi-version", "v", "", "Manually defines a GLPI version (Format: x.x.x)")
}
