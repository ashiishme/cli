package kafka

import (
	"sort"

	"github.com/antihax/optional"
	"github.com/confluentinc/go-printer"
	"github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
	"github.com/spf13/cobra"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/examples"
	"github.com/confluentinc/cli/internal/pkg/output"
	"github.com/confluentinc/cli/internal/pkg/properties"
	"github.com/confluentinc/cli/internal/pkg/utils"
)

func (c *brokerCommand) newUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [id]",
		Args:  cobra.MaximumNArgs(1),
		RunE:  c.update,
		Short: "Update per-broker or cluster-wide Kafka broker configs.",
		Example: examples.BuildExampleString(
			examples.Example{
				Text: "Update configuration values for broker 1.",
				Code: "confluent kafka broker update 1 --config min.insync.replicas=2,num.partitions=2",
			},
			examples.Example{
				Text: "Update configuration values for all brokers in the cluster.",
				Code: "confluent kafka broker update --all --config min.insync.replicas=2,num.partitions=2",
			},
		),
	}

	cmd.Flags().StringSlice("config", nil, `A comma-separated list of configuration overrides ("key=value") for the broker being updated.`)
	cmd.Flags().Bool("all", false, "Apply config update to all brokers in the cluster.")
	cmd.Flags().AddFlagSet(pcmd.OnPremKafkaRestSet())
	pcmd.AddOutputFlag(cmd)

	_ = cmd.MarkFlagRequired("config")

	return cmd
}

func (c *brokerCommand) update(cmd *cobra.Command, args []string) error {
	brokerId, all, err := checkAllOrBrokerIdSpecified(cmd, args)
	if err != nil {
		return err
	}

	format, err := cmd.Flags().GetString(output.FlagName)
	if err != nil {
		return err
	}

	restClient, restContext, err := initKafkaRest(c.AuthenticatedCLICommand, cmd)
	if err != nil {
		return err
	}

	clusterId, err := getClusterIdForRestRequests(restClient, restContext)
	if err != nil {
		return err
	}

	configs, err := cmd.Flags().GetStringSlice("config")
	if err != nil {
		return err
	}
	configMap, err := properties.ConfigFlagToMap(configs)
	if err != nil {
		return err
	}
	data := toAlterConfigBatchRequestData(configMap)

	if all {
		resp, err := restClient.ConfigsV3Api.UpdateKafkaClusterConfigs(restContext, clusterId,
			&kafkarestv3.UpdateKafkaClusterConfigsOpts{
				AlterConfigBatchRequestData: optional.NewInterface(data),
			})
		if err != nil {
			return kafkaRestError(restClient.GetConfig().BasePath, err, resp)
		}
	} else {
		resp, err := restClient.ConfigsV3Api.ClustersClusterIdBrokersBrokerIdConfigsalterPost(restContext, clusterId, brokerId,
			&kafkarestv3.ClustersClusterIdBrokersBrokerIdConfigsalterPostOpts{
				AlterConfigBatchRequestData: optional.NewInterface(data),
			})
		if err != nil {
			return kafkaRestError(restClient.GetConfig().BasePath, err, resp)
		}
	}

	if format == output.Human.String() {
		c.printHumanUpdate(all, clusterId, brokerId, data)
		return nil
	}

	return c.printStructuredUpdate(format, data)
}

func (c *brokerCommand) printHumanUpdate(all bool, clusterId string, brokerId int32, data kafkarestv3.AlterConfigBatchRequestData) {
	if all {
		utils.Printf(c.Command, "Updated the following broker configs for cluster \"%s\":\n", clusterId)
	} else {
		utils.Printf(c.Command, "Updated the following configs for broker \"%d\":\n", brokerId)
	}
	tableLabels := []string{"Name", "Value"}
	tableEntries := make([][]string, len(data.Data))
	for i, config := range data.Data {
		tableEntries[i] = printer.ToRow(
			&struct {
				Name  string
				Value string
			}{Name: config.Name, Value: *config.Value}, []string{"Name", "Value"})
	}
	sort.Slice(tableEntries, func(i, j int) bool {
		return tableEntries[i][0] < tableEntries[j][0]
	})
	printer.RenderCollectionTable(tableEntries, tableLabels)
}

func (c *brokerCommand) printStructuredUpdate(format string, data kafkarestv3.AlterConfigBatchRequestData) error {
	type printConfig struct {
		Name  string `json:"name" yaml:"name"`
		Value string `json:"value,omitempty" yaml:"value,omitempty"`
	}
	printConfigs := make([]*printConfig, len(data.Data))
	for i, config := range data.Data {
		printConfigs[i] = &printConfig{
			Name:  config.Name,
			Value: *config.Value,
		}
	}
	sort.Slice(printConfigs, func(i, j int) bool {
		return printConfigs[i].Name < printConfigs[j].Name
	})
	return output.StructuredOutput(format, printConfigs)
}
