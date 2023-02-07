package cmd

import (
	"fmt"

	"github.com/Feals-404/GLPIAnarchy/exploits"
	"github.com/spf13/cobra"
)

var sleep *int
var minfolder *int
var maxfolder *int
var minfile *int
var maxfile *int
var numRoutines *int

var xploitCmd = &cobra.Command{
	Use:     "xploit [url] [exploit]",
	Short:   "Exploit the GLPI with a specified vulnerability",
	Example: `GLPIAnarchy xploit https://glpi.example.com CVE-2020-15175 [Optional flags]`,
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		target := exploits.NormalizeURL(args[0])
		exploit := args[1]

		for _, vulnerability := range exploits.Vulnerabilities {
			if exploit == vulnerability.ID {
				fmt.Print("\nExploiting " + target + " with " + exploit + "\n\n")
				if vulnerability.TAG == "TB SQLi" {
					vulnerability.RUNTBSQL(target, *sleep)
					break
				} else if vulnerability.ID == "CVE-2023-22500" {
					exploits.CVE202322500(target, *numRoutines, *minfolder, *maxfolder, *minfile, *maxfile)

				} else {
					vulnerability.RUN(target)
				}

			}
		}

	},
}

func init() {
	rootCmd.AddCommand(xploitCmd)

	sleep = xploitCmd.Flags().IntP("sleep", "s", 5, "Sleep time for TB SQLi")
	minfolder = xploitCmd.Flags().IntP("minfolder", "I", 0, "Min folder to start (For CVE-2023-22500)")
	maxfolder = xploitCmd.Flags().IntP("maxfolder", "A", 50, "Max folder (For CVE-2023-22500)")
	minfile = xploitCmd.Flags().IntP("filemin", "i", 0, "Min file to start (For CVE-2023-22500)")
	maxfile = xploitCmd.Flags().IntP("filemax", "a", 10, "Max file (For CVE-2023-22500)")
	numRoutines = xploitCmd.Flags().IntP("numroutine", "r", 5, "Number of go routines (For CVE-2023-22500)")

}
