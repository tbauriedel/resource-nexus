package tf

import (
	"context"
	"os"
	"os/exec"
)

type SubCommand string

const (
	SubCommandInit  SubCommand = "init"
	SubCommandPlan  SubCommand = "plan"
	SubCommandApply SubCommand = "apply"
)

type Command struct {
	*exec.Cmd
}

// GetCommandInit returns a new Command instance for terraform init.
func (tf *TerraformInstance) GetCommandInit() *Command {
	return tf.buildCommand(SubCommandInit, nil)
}

// GetCommandPlan returns a new Command instance for terraform plan.
func (tf *TerraformInstance) GetCommandPlan() *Command {
	return tf.buildCommand(SubCommandPlan, nil)
}

// GetCommandApply returns a new Command instance for terraform apply.
func (tf *TerraformInstance) GetCommandApply() *Command {
	return tf.buildCommand(SubCommandApply, nil)
}

// buildCommand builds a new Command instance.
// tmpWorkDir will be added as the working directory.
// subcommand and args will be appended to the command arguments.
func (tf *TerraformInstance) buildCommand(subcommand SubCommand, args []string) *Command {
	c := Command{}

	// Build full argument chain
	fullArgs := append([]string{string(subcommand)}, args...)
	fullArgs = append(fullArgs, "--json")

	// Build command
	ctx, cancel := context.WithTimeout(context.Background(), tf.CommandTimeout)
	defer cancel()

	c.Cmd = exec.CommandContext(ctx, tf.ExecutablePath, fullArgs...) //nolint:gosec

	// add tmp working directory
	c.Dir = tf.tmpWorkDir

	c.Cancel = func() error {
		return c.Process.Signal(os.Interrupt)
	}

	// Prepare command arguments

	return &c
}
