package cmd

import (
	"fmt"

	"github.com/Feals-404/GLPIAnarchy/exploits"
	"github.com/spf13/cobra"
)

var sleep string
var xploitCmd = &cobra.Command{
	Use:     "xploit [url] [exploit]",
	Short:   "Exploit the GLPI with a specified vulnerability",
	Example: `GLPIAnarchy xploit https://glpi.example.com CVE-2020-15175`,
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		target := exploits.NormalizeURL(args[0])
		exploit := args[1]

		for _, vulnerability := range exploits.Vulnerabilities {
			if exploit == vulnerability.ID {
				fmt.Print("\nExploiting " + target + " with " + exploit + "\n\n")
				if vulnerability.TAG == "TB SQLi" {
					vulnerability.RUNTBSQL(target, sleep)
					break
				}
				vulnerability.RUN(target)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(xploitCmd)
	xploitCmd.Flags().StringVarP(&sleep, "sleep", "s", "", "Sleep time for TB SQLi")
}
