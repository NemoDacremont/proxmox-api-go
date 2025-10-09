package create

import (
	"context"

	"github.com/Telmate/proxmox-api-go/cli"
	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

var create_sdnzoneCmd = &cobra.Command{
	Use:   "sdnzone",
	Short: "Creates a new SDN zone with id SDNZONE and type SDNTYPE",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return createSDN(cli.Context(), args)
	},
}

func init() {
	CreateCmd.AddCommand(create_sdnzoneCmd)
}

func createSDN(ctx context.Context, args []string) (err error) {
	client := cli.NewClient()

	config, err := proxmox.NewConfigSDNZoneFromJson(cli.NewConfig())
	if err != nil {
		return err
	}

	if err := config.CreateWithValidate(ctx, config.Zone, client); err != nil {
		return err
	}

	cli.PrintItemCreated(CreateCmd.OutOrStdout(), config.Zone, "SDN_" + config.Type)
	return nil
}
