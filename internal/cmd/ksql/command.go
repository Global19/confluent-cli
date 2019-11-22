package ksql

import (
	"github.com/spf13/cobra"

	"github.com/confluentinc/ccloud-sdk-go"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/config"
)

type command struct {
	*cobra.Command
	config      *config.Config
	client      ccloud.KSQL
	kafkaClient ccloud.Kafka
	userClient  ccloud.User
	ch          *pcmd.ConfigHelper
}

// New returns the default command object for interacting with KSQL.
func New(prerunner pcmd.PreRunner, config *config.Config, client ccloud.KSQL,
	kafkaClient ccloud.Kafka, userClient ccloud.User, ch *pcmd.ConfigHelper) *cobra.Command {
	cmd := &command{
		Command: &cobra.Command{
			Use:               "ksql",
			Short:             "Manage KSQL applications.",
			PersistentPreRunE: prerunner.Authenticated(),
		},
		config:      config,
		client:      client,
		kafkaClient: kafkaClient,
		userClient:  userClient,
		ch:          ch,
	}
	cmd.init()
	return cmd.Command
}

func (c *command) init() {
	c.AddCommand(NewClusterCommand(c.config, c.client, c.kafkaClient, c.userClient, c.ch))
}
