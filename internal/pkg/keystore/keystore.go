//go:generate go run github.com/travisjeffery/mocker/cmd/mocker --prefix "" --dst ../mock/keystore.go --pkg mock keystore.go KeyStore
package keystore

import (
	schedv1 "github.com/confluentinc/cc-structs/kafka/scheduler/v1"
	"github.com/spf13/cobra"

	"github.com/confluentinc/cli/internal/pkg/cmd"
	v0 "github.com/confluentinc/cli/internal/pkg/config/v0"
	"github.com/confluentinc/cli/internal/pkg/errors"
)

type KeyStore interface {
	HasAPIKey(key string, clusterId string, cmd *cobra.Command) (bool, error)
	StoreAPIKey(key *schedv1.ApiKey, clusterId string, cmd *cobra.Command) error
	DeleteAPIKey(key string) error
}

type ConfigKeyStore struct {
	Config *cmd.DynamicConfig
}

func (c *ConfigKeyStore) HasAPIKey(key string, clusterId string, cmd *cobra.Command) (bool, error) {
	ctx := c.Config.Context()
	if ctx == nil {
		return false, new(errors.NotLoggedInError)
	}
	kcc, err := ctx.FindKafkaCluster(cmd, clusterId)
	if err != nil {
		return false, err
	}
	_, found := kcc.APIKeys[key]
	return found, nil
}

// StoreAPIKey creates a new API key pair in the local key store for later usage
func (c *ConfigKeyStore) StoreAPIKey(key *schedv1.ApiKey, clusterId string, cmd *cobra.Command) error {
	ctx := c.Config.Context()
	if ctx == nil {
		return new(errors.NotLoggedInError)
	}
	kcc, err := ctx.FindKafkaCluster(cmd, clusterId)
	if err != nil {
		return err
	}
	kcc.APIKeys[key.Key] = &v0.APIKeyPair{
		Key:    key.Key,
		Secret: key.Secret,
	}
	return c.Config.Save()
}

func (c *ConfigKeyStore) DeleteAPIKey(key string) error {
	ctx := c.Config.Context()
	if ctx == nil {
		return new(errors.NotLoggedInError)
	}
	ctx.KafkaClusterContext.DeleteAPIKey(key)
	return c.Config.Save()
}
