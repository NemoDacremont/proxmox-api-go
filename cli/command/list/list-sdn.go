package list

import (
	"github.com/Telmate/proxmox-api-go/cli"
	"github.com/spf13/cobra"
)

var list_sdnCmd = &cobra.Command{
	Use:   "sdnzones",
	Short: "Prints a list of SDN Zone in raw json format",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		listRaw(cli.Context(), "SdnZones")
	},
}

func init() {
	listCmd.AddCommand(list_sdnCmd)
}
