package streamshare

import (
	"github.com/spf13/cobra"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/examples"
	"github.com/confluentinc/cli/internal/pkg/output"
)

func (c *command) newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List shares for provider.",
		Args:  cobra.NoArgs,
		RunE:  c.list,
		Example: examples.BuildExampleString(
			examples.Example{
				Text: "List provider shares:",
				Code: "confluent stream-share provider share list",
			},
		),
	}

	cmd.Flags().String("shared-resource", "", "Filter the results by exact match for shared resource.")

	pcmd.AddOutputFlag(cmd)

	return cmd
}

func (c *command) list(cmd *cobra.Command, _ []string) error {
	sharedResource, err := cmd.Flags().GetString("shared-resource")
	if err != nil {
		return err
	}

	providerShares, err := c.V2Client.ListProviderShares(sharedResource)
	if err != nil {
		return err
	}

	outputWriter, err := output.NewListOutputWriter(cmd, providerShareListFields, providerShareListHumanLabels, providerShareListStructuredLabels)
	if err != nil {
		return err
	}

	for _, share := range providerShares {
		element := c.buildProviderShare(share)
		outputWriter.AddElement(element)
	}

	return outputWriter.Out()
}