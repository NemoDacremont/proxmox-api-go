package delete

import (
	"github.com/Telmate/proxmox-api-go/cli"
	"github.com/spf13/cobra"
)

var delete_sdnzoneCmd = &cobra.Command{
	Use:   "sdnzone SDNZONEID",
	Short: "Deletes the SDN Zone with id SDNZONE",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return deleteID(cli.Context(), args, "SdnZone")
	},
}

func init() {
	deleteCmd.AddCommand(delete_sdnzoneCmd)
}
