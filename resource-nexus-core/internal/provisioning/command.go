package provisioning

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

type Command struct {
	*exec.Cmd
}

type SubCommand string

const (
	SubCommandInit  SubCommand = "init"
	SubCommandPlan  SubCommand = "plan"
	SubCommandApply SubCommand = "apply"
)

// GetCommandInit returns the command for `<provisioner> init`.
//
// The command is built with the given arguments and context.
// Arguments are appended to the command string.
func (bp *BaseProvisioner) GetCommandInit(ctx context.Context, args []string) (*Command, error) {
	err := bp.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid provisioner settings: %w", err)
	}

	return buildCommand(
		bp.WorkingDirectory,
		bp.ExecutablePath,
		SubCommandInit,
		args,
		ctx,
	), nil
}

// GetCommandPlan returns the command for `<provisioner> plan`.
//
// The command is built with the given arguments and context.
// Arguments are appended to the command string.
func (bp *BaseProvisioner) GetCommandPlan(ctx context.Context, args []string) (*Command, error) {
	err := bp.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid provisioner settings: %w", err)
	}

	return buildCommand(
		bp.WorkingDirectory,
		bp.ExecutablePath,
		SubCommandPlan,
		args,
		ctx,
	), nil
}

// GetCommandApply returns the command for `<provisioner> apply`.
//
// The command is built with the given arguments and context.
// Arguments are appended to the command string.
func (bp *BaseProvisioner) GetCommandApply(ctx context.Context, args []string) (*Command, error) {
	err := bp.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid provisioner settings: %w", err)
	}

	return buildCommand(
		bp.WorkingDirectory,
		bp.ExecutablePath,
		SubCommandApply,
		args,
		ctx,
	), nil
}

// buildCommand returns a new Command.
//
// workdir, executable and subcommand are used to build the command string.
// ctx that will be used for the Command.
//
// Ensure the provided executable, subcommand and args are validated and no injection attacks are possible!
func buildCommand(
	workdir string, executable string, subcommand SubCommand, args []string, ctx context.Context,
) *Command {
	command := Command{}

	// add subcommand and custom arguments to the string array
	arguments := append([]string{string(subcommand)}, args...)
	// append json argument
	arguments = append(arguments, "--json")

	// builds command with context
	// The allowed exec commands are defined inside the config files.
	// They are validated inside the command functions above!
	command.Cmd = exec.CommandContext(ctx, executable, arguments...) //nolint:gosec

	// specify work directory
	command.Dir = workdir

	// define the command that will be triggered once the context is reached / canceled
	command.Cancel = func() error {
		return command.Process.Signal(os.Interrupt)
	}

	return &command
}
