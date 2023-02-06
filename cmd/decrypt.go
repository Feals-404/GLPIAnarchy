package cmd

import (
	"fmt"

	"github.com/Feals-404/GLPIAnarchy/exploits"
	"github.com/spf13/cobra"
)

var key string
var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Short:   "Decrypt retrieved Passwd string. The target must be below version 9.4.6",
	Example: "GLPIAnarchy decrypt [string]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(exploits.DecryptLDAP(args[0], key))
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringVarP(&key, "key", "k", "", "Non-Default key")
}
