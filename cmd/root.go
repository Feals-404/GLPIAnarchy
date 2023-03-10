package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GLPIAnarchy",
	Short: "Multi Exploit GLPI",
	Args:  cobra.ExactArgs(1),
	Long: `

▄█     ▄███████▄  ▄█          ▄██████▄          ▄████████ ███▄▄▄▄      ▄████████    ▄████████  ▄████████    ▄█    █▄    ▄██   ▄   
███    ███    ███ ███         ███    ███        ███    ███ ███▀▀▀██▄   ███    ███   ███    ███ ███    ███   ███    ███   ███   ██▄ 
███▌   ███    ███ ███         ███    █▀         ███    ███ ███   ███   ███    ███   ███    ███ ███    █▀    ███    ███   ███▄▄▄███ 
███▌   ███    ███ ███        ▄███               ███    ███ ███   ███   ███    ███  ▄███▄▄▄▄██▀ ███         ▄███▄▄▄▄███▄▄ ▀▀▀▀▀▀███ 
███▌ ▀█████████▀  ███       ▀▀███ ████▄       ▀███████████ ███   ███ ▀███████████ ▀▀███▀▀▀▀▀   ███        ▀▀███▀▀▀▀███▀  ▄██   ███ 
███    ███        ███         ███    ███        ███    ███ ███   ███   ███    ███ ▀███████████ ███    █▄    ███    ███   ███   ███ 
███    ███        ███▌    ▄   ███    ███        ███    ███ ███   ███   ███    ███   ███    ███ ███    ███   ███    ███   ███   ███ 
█▀    ▄████▀      █████▄▄██   ████████▀         ███    █▀   ▀█   █▀    ███    █▀    ███    ███ ████████▀    ███    █▀     ▀█████▀  
                                  ▀                                                                ███     ███                                     
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
