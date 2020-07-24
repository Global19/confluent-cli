package config

import (
	"sort"

	"github.com/confluentinc/go-printer"
	"github.com/spf13/cobra"

	"github.com/confluentinc/cli/internal/pkg/analytics"
	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/errors"
	"github.com/confluentinc/cli/internal/pkg/output"
)

var (
	contextListFields           = []string{"Current", "Name", "Platform", "Credential"}
	contextListHumanLabels      = []string{"Current", "Name", "Platform", "Credential"}
	contextListStructuredLabels = []string{"current", "name", "platform", "credential"}
)

type contextCommand struct {
	*pcmd.CLICommand
	prerunner pcmd.PreRunner
	analytics analytics.Client
}

// NewContext returns the Cobra contextCommand for `config context`.
func NewContext(prerunner pcmd.PreRunner, analytics analytics.Client) *cobra.Command {
	cliCmd := pcmd.NewAnonymousCLICommand(
		&cobra.Command{
			Use:   "context",
			Short: "Manage config contexts.",
		}, prerunner)
	cmd := &contextCommand{
		CLICommand: cliCmd,
		prerunner:  prerunner,
		analytics:  analytics,
	}
	cmd.init()
	return cmd.Command
}

func (c *contextCommand) init() {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all config contexts.",
		RunE:  pcmd.NewCLIRunE(c.list),
		Args:  cobra.NoArgs,
	}
	listCmd.Flags().StringP(output.FlagName, output.ShortHandFlag, output.DefaultValue, output.Usage)
	listCmd.Flags().SortFlags = false
	c.AddCommand(listCmd)
	c.AddCommand(&cobra.Command{
		Use:   "use <id>",
		Short: "Use a config context.",
		RunE:  pcmd.NewCLIRunE(c.use),
		Args:  cobra.ExactArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			c.analytics.SetCommandType(analytics.ContextUse)
			return c.prerunner.Anonymous(c.CLICommand)(cmd, args)
		},
	})
	c.AddCommand(&cobra.Command{
		Use:   "current",
		Short: "Show the current config context.",
		RunE:  pcmd.NewCLIRunE(c.current),
		Args:  cobra.NoArgs,
	})

	getCmd := &cobra.Command{
		Use:   "get <id or no argument for current context>",
		Short: "Get a config context parameter.",
		RunE:  pcmd.NewCLIRunE(c.get),
		Args:  cobra.RangeArgs(0, 1),
	}
	getCmd.Hidden = true
	c.AddCommand(getCmd)

	setCmd := &cobra.Command{
		Use:   "set <id or no argument for current context>",
		Short: "Set a config context parameter.",
		RunE:  pcmd.NewCLIRunE(c.set),
		Args:  cobra.RangeArgs(0, 1),
	}
	setCmd.Flags().String("kafka-cluster", "", "Set the current Kafka cluster context.")
	setCmd.Flags().SortFlags = false
	setCmd.Hidden = true
	c.AddCommand(setCmd)

	c.AddCommand(&cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a config context.",
		RunE:  pcmd.NewCLIRunE(c.delete),
		Args:  cobra.ExactArgs(1),
	})
}

func (c *contextCommand) list(cmd *cobra.Command, _ []string) error {
	type row struct {
		Current    string
		Name       string
		Platform   string
		Credential string
	}
	var contextNames []string
	for name := range c.Config.Contexts {
		contextNames = append(contextNames, name)
	}
	sort.Strings(contextNames)
	outputWriter, err := output.NewListOutputWriter(cmd, contextListFields, contextListHumanLabels, contextListStructuredLabels)
	if err != nil {
		return err
	}
	for _, name := range contextNames {
		context := c.Config.Contexts[name]
		current := ""
		// Add '*' only in the case where we are printing out tables
		if outputWriter.GetOutputFormat() == output.Human {
			if c.Config.CurrentContext == name {
				current = "*"
			}
		} else {
			current = "false"
			if c.Config.CurrentContext == name {
				current = "true"
			}
		}
		outputWriter.AddElement(&row{current, name, context.PlatformName, context.CredentialName})
	}
	return outputWriter.Out()
}

func (c *contextCommand) use(cmd *cobra.Command, args []string) error {
	name := args[0]
	err := c.Config.SetContext(name)
	if err != nil {
		return err
	}
	return nil
}

func (c *contextCommand) current(cmd *cobra.Command, _ []string) error {
	pcmd.Println(cmd, c.Config.CurrentContext)
	return nil
}

func (c *contextCommand) get(cmd *cobra.Command, args []string) error {
	context, err := c.context(cmd, args)
	if err != nil {
		return err
	}
	return printer.RenderYAMLOut(context, nil, nil, cmd.OutOrStdout())
}

func (c *contextCommand) set(cmd *cobra.Command, args []string) error {
	context, err := c.context(cmd, args)
	if err != nil {
		return err
	}
	if cmd.Flags().Changed("kafka-cluster") {
		clusterId, err := cmd.Flags().GetString("kafka-cluster")
		if err != nil {
			return err
		}
		return context.SetActiveKafkaCluster(cmd, clusterId)
	}
	return nil
}

func (c *contextCommand) delete(cmd *cobra.Command, args []string) error {
	contextName := args[0]
	err := c.Config.DeleteContext(contextName)
	if err != nil {
		return err
	}
	return c.Config.Save()
}

func (c *contextCommand) context(cmd *cobra.Command, args []string) (*pcmd.DynamicContext, error) {
	var context *pcmd.DynamicContext
	var err error
	if len(args) == 1 {
		contextName := args[0]
		context, err = c.Config.FindContext(contextName)
	} else {
		context, err = c.Config.Context(cmd)
		if context == nil {
			err = &errors.NoContextError{CLIName: c.Config.CLIName}
		}
	}
	if err != nil {
		return nil, err
	}
	return context, nil
}
