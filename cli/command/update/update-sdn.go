package update

import (
	"github.com/Telmate/proxmox-api-go/cli"
	"github.com/spf13/cobra"
)

var update_sdnCmd = &cobra.Command{
	Use:   "sdn",
	Short: "Apply SDN configuration",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		client := cli.NewClient()
		if _, err := client.ApplySDN(cli.Context()); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	updateCmd.AddCommand(update_sdnCmd)
}
