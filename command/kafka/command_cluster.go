package kafka

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"

	chttp "github.com/confluentinc/ccloud-sdk-go"
	authv1 "github.com/confluentinc/ccloudapis/auth/v1"
	kafkav1 "github.com/confluentinc/ccloudapis/kafka/v1"
	"github.com/confluentinc/cli/command/common"
	"github.com/confluentinc/cli/shared"
	"github.com/confluentinc/go-printer"
)

var (
	listFields      = []string{"Id", "Name", "ServiceProvider", "Region", "Durability", "Status"}
	listLabels      = []string{"Id", "Name", "Provider", "Region", "Durability", "Status"}
	describeFields  = []string{"Id", "Name", "NetworkIngress", "NetworkEgress", "Storage", "ServiceProvider", "Region", "Status", "Endpoint", "ApiEndpoint", "PricePerHour"}
	describeRenames = map[string]string{"NetworkIngress": "Ingress", "NetworkEgress": "Egress", "ServiceProvider": "Provider"}
)

type clusterCommand struct {
	*cobra.Command
	config *shared.Config
	client chttp.Kafka
}

// NewClusterCommand returns the Cobra clusterCommand for Kafka Cluster.
func NewClusterCommand(config *shared.Config, plugin common.GRPCPlugin) *cobra.Command {
	cmd := &clusterCommand{
		Command: &cobra.Command{
			Use:   "cluster",
			Short: "Manage Kafka clusters",
		},
		config: config,
	}
	cmd.init(plugin)
	return cmd.Command
}

func (c *clusterCommand) init(plugin common.GRPCPlugin) {
	c.Command.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := common.SetLoggingVerbosity(cmd, c.config.Logger); err != nil {
			return common.HandleError(err, cmd)
		}
		if err := c.config.CheckLogin(); err != nil {
			return common.HandleError(err, cmd)
		}
		// Lazy load plugin to avoid unnecessarily spawning child processes
		return plugin.Load(&c.client, c.config.Logger)
	}

	c.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List Kafka clusters",
		RunE:  c.list,
		Args:  cobra.NoArgs,
	})

	createCmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Create a Kafka cluster",
		RunE:  c.create,
		Args:  cobra.ExactArgs(1),
	}
	createCmd.Flags().String("cloud", "", "aws or gcp")
	createCmd.Flags().String("region", "", "a valid region in the given cloud")
	// default to smallest size allowed
	createCmd.Flags().Int32("ingress", 1, "network ingress in MB/s")
	createCmd.Flags().Int32("egress", 1, "network egress in MB/s")
	createCmd.Flags().Int32("storage", 500, "total usable data storage in GB")
	createCmd.Flags().Bool("multizone", false, "use multiple zones for high availability")
	check(createCmd.MarkFlagRequired("cloud"))
	check(createCmd.MarkFlagRequired("region"))
	c.AddCommand(createCmd)

	c.AddCommand(&cobra.Command{
		Use:   "describe ID",
		Short: "Describe a Kafka cluster",
		RunE:  c.describe,
		Args:  cobra.ExactArgs(1),
	})

	updateCmd := &cobra.Command{
		Use:   "update ID",
		Short: "Update a Kafka cluster",
		RunE:  c.update,
		Args:  cobra.ExactArgs(1),
	}
	updateCmd.Hidden = true
	c.AddCommand(updateCmd)

	c.AddCommand(&cobra.Command{
		Use:   "delete ID",
		Short: "Delete a Kafka cluster",
		RunE:  c.delete,
		Args:  cobra.ExactArgs(1),
	})
	c.AddCommand(&cobra.Command{
		Use:   "auth",
		Short: "Configure authorization for a Kafka cluster",
		RunE:  c.auth,
		Args:  cobra.NoArgs,
	})
	c.AddCommand(&cobra.Command{
		Use:   "use ID",
		Short: "Make the Kafka cluster active for use in other commands",
		RunE:  c.use,
		Args:  cobra.ExactArgs(1),
	})
}

func (c *clusterCommand) list(cmd *cobra.Command, args []string) error {
	req := &kafkav1.KafkaCluster{AccountId: c.config.Auth.Account.Id}
	clusters, err := c.client.List(context.Background(), req)
	if err != nil {
		return common.HandleError(err, cmd)
	}
	currCtx, err := c.config.Context()
	if err != nil && err != shared.ErrNoContext {
		return err
	}
	var data [][]string
	for _, cluster := range clusters {
		if cluster.Id == currCtx.Kafka {
			cluster.Id = fmt.Sprintf("* %s", cluster.Id)
		} else {
			cluster.Id = fmt.Sprintf("  %s", cluster.Id)
		}
		data = append(data, printer.ToRow(cluster, listFields))
	}
	printer.RenderCollectionTable(data, listLabels)
	return nil
}

func (c *clusterCommand) create(cmd *cobra.Command, args []string) error {
	cloud, err := cmd.Flags().GetString("cloud")
	if err != nil {
		return common.HandleError(err, cmd)
	}
	region, err := cmd.Flags().GetString("region")
	if err != nil {
		return common.HandleError(err, cmd)
	}
	ingress, err := cmd.Flags().GetInt32("ingress")
	if err != nil {
		return common.HandleError(err, cmd)
	}
	egress, err := cmd.Flags().GetInt32("egress")
	if err != nil {
		return common.HandleError(err, cmd)
	}
	storage, err := cmd.Flags().GetInt32("storage")
	if err != nil {
		return common.HandleError(err, cmd)
	}
	multizone, err := cmd.Flags().GetBool("multizone")
	if err != nil {
		return common.HandleError(err, cmd)
	}
	durability := kafkav1.Durability_LOW
	if multizone {
		durability = kafkav1.Durability_HIGH
	}
	config := &kafkav1.KafkaClusterConfig{
		AccountId:       c.config.Auth.Account.Id,
		Name:            args[0],
		ServiceProvider: cloud,
		Region:          region,
		NetworkIngress:  ingress,
		NetworkEgress:   egress,
		Storage:         storage,
		Durability:      durability,
	}
	cluster, err := c.client.Create(context.Background(), config)
	if err != nil {
		// TODO: don't swallow validation errors (reportedly separately)
		return common.HandleError(err, cmd)
	}
	return printer.RenderTableOut(cluster, describeFields, describeRenames, os.Stdout)
}

func (c *clusterCommand) describe(cmd *cobra.Command, args []string) error {
	req := &kafkav1.KafkaCluster{AccountId: c.config.Auth.Account.Id, Id: args[0]}
	cluster, err := c.client.Describe(context.Background(), req)
	if err != nil {
		return common.HandleError(err, cmd)
	}
	return printer.RenderTableOut(cluster, describeFields, describeRenames, os.Stdout)
}

func (c *clusterCommand) update(cmd *cobra.Command, args []string) error {
	return shared.ErrNotImplemented
}

func (c *clusterCommand) delete(cmd *cobra.Command, args []string) error {
	req := &kafkav1.KafkaCluster{AccountId: c.config.Auth.Account.Id, Id: args[0]}
	err := c.client.Delete(context.Background(), req)
	if err != nil {
		return common.HandleError(err, cmd)
	}
	fmt.Printf("The Kafka cluster %s has been deleted.\n", args[0])
	return nil
}

func (c *clusterCommand) auth(cmd *cobra.Command, args []string) error {
	cfg, err := c.config.Context()

	if err != nil {
		return common.HandleError(err, cmd)
	}

	if cfg.Kafka == "" {
		return fmt.Errorf("No cluster selected. See ccloud kafka use for help ")
	}

	cluster, known := c.config.Platforms[cfg.Platform].KafkaClusters[cfg.Kafka]
	if known {
		fmt.Printf("Kafka Cluster: %s\n", cfg.Kafka)
		fmt.Printf("Bootstrap Servers: %s\n", cluster.Bootstrap)
		fmt.Printf("API Key: %s\n", cluster.APIKey)
		fmt.Printf("API Secret: %s\n", cluster.APISecret)
		return nil
	}

	userProvidingKey, err := userHasKey(cfg.Kafka)
	if err != nil {
		return common.HandleError(err, cmd)
	}

	var key, secret string
	if userProvidingKey {
		key, secret, err = promptForKafkaCreds()
	} else {
		key, secret, err = c.createKafkaCreds(context.Background(), cfg.Kafka)
	}
	if err != nil {
		return common.HandleError(err, cmd)
	}

	req := &kafkav1.KafkaCluster{AccountId: c.config.Auth.Account.Id, Id: cfg.Kafka}
	kc, err := c.client.Describe(context.Background(), req)
	if err != nil {
		return common.HandleError(err, cmd)
	}

	if c.config.Platforms[cfg.Platform].KafkaClusters == nil {
		c.config.Platforms[cfg.Platform].KafkaClusters = map[string]shared.KafkaClusterConfig{}
	}
	c.config.Platforms[cfg.Platform].KafkaClusters[cfg.Kafka] = shared.KafkaClusterConfig{
		Bootstrap:   strings.TrimPrefix(kc.Endpoint, "SASL_SSL://"),
		APIEndpoint: kc.ApiEndpoint,
		APIKey:      key,
		APISecret:   secret,
	}
	return c.config.Save()
}

func (c *clusterCommand) use(cmd *cobra.Command, args []string) error {
	cfg, err := c.config.Context()
	if err != nil {
		return common.HandleError(err, cmd)
	}
	cfg.Kafka = args[0]
	return c.config.Save()
}

//
// Helper functions
//

func userHasKey(kafkaClusterID string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Do you have an API key for %s? [N/y] ", kafkaClusterID)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	r := strings.TrimSpace(response)
	return r == "" || r[0] == 'y' || r[0] == 'Y', nil
}

func promptForKafkaCreds() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("API Key: ")
	key, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("API Secret: ")
	byteSecret, err := terminal.ReadPassword(0)
	fmt.Println()
	if err != nil {
		return "", "", err
	}
	secret := string(byteSecret)

	return strings.TrimSpace(key), strings.TrimSpace(secret), nil
}

func (c *clusterCommand) createKafkaCreds(ctx context.Context, kafkaClusterID string) (string, string, error) {
	client := chttp.NewClientWithJWT(ctx, c.config.AuthToken, c.config.AuthURL, c.config.Logger)
	key, err := client.APIKey.Create(ctx, &authv1.ApiKey{
		UserId: c.config.Auth.User.Id,
		LogicalClusters: []*authv1.ApiKey_Cluster{
			{Id: kafkaClusterID},
		},
		AccountId: c.config.Auth.Account.Id,
	})

	if err != nil {
		return "", "", shared.ConvertAPIError(err)
	}
	fmt.Println("Okay, we've created an API key. If needed, you can see it with `ccloud kafka cluster auth`.")
	return key.Key, key.Secret, nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
