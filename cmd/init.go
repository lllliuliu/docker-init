// The "init" subcommand contains the follow logic:
// - Search for unuesd ports from 10000 to 65535 in order,
//   assign exposed port to different services and write to .env file
// - Use the above port to build to name to generate the network
package cmd

import (
	"fmt"

	"docker-init/core"

	"github.com/spf13/cobra"
)

// used for flags
var composerYML string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize .env file and generate the network",
	Long: `Search for unuesd ports from 10000 to 65535 in order, assign exposed port to 
different services and write to .env file. Use the above port to build to name 
to generate the network.`,
	Run: func(cmd *cobra.Command, args []string) {
		message := core.Cinit(composerYML)
		fmt.Println(message)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// cc represent to file path of docker-compose.yml, default current directory
	initCmd.Flags().StringVar(&composerYML, "cc", "./docker-compose.yml", "Specify the docker-compose.yml file path")
}
