package iam

import (
	"github.com/spf13/cobra"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	v3 "github.com/confluentinc/cli/internal/pkg/config/v3"
	"github.com/confluentinc/cli/internal/pkg/shell/completer"
)

type command struct {
	*pcmd.AuthenticatedCLICommand
	prerunner pcmd.PreRunner
}

// New returns the default command object for interacting with RBAC.
func New(cfg *v3.Config, prerunner pcmd.PreRunner, serverCompleter completer.ServerSideCompleter) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "iam",
		Annotations: map[string]string{pcmd.RunRequirement: pcmd.RequireNonAPIKeyCloudLoginOrOnPremLogin},
	}

	var cliCmd *pcmd.AuthenticatedCLICommand
	if cfg.IsOnPremLogin() {
		cmd.Short = "Manage RBAC, ACL and IAM permissions."
		cmd.Long = "Manage Role-Based Access Control (RBAC), Access Control Lists (ACL), and Identity and Access Management (IAM) permissions."
		cliCmd = pcmd.NewAuthenticatedWithMDSCLICommand(cmd, prerunner)
	} else {
		cmd.Short = "Manage RBAC and IAM permissions."
		cmd.Long = "Manage Role-Based Access Control (RBAC) and Identity and Access Management (IAM) permissions."
		cliCmd = pcmd.NewAuthenticatedCLICommand(cmd, prerunner)
	}

	c := &command{
		AuthenticatedCLICommand: cliCmd,
		prerunner:               prerunner,
	}

	serviceAccountCmd := NewServiceAccountCommand(c.prerunner)

	c.AddCommand(NewACLCommand(c.prerunner))
	c.AddCommand(NewRoleCommand(cfg, c.prerunner))
	c.AddCommand(NewRolebindingCommand(cfg, c.prerunner))
	c.AddCommand(serviceAccountCmd.Command)
	c.AddCommand(NewUserCommand(c.prerunner))

	if cfg.IsCloudLogin() {
		serverCompleter.AddCommand(serviceAccountCmd)
	}

	return c.Command
}
