package provisioning

import (
	"context"
	"testing"

	"github.com/tbauriedel/resource-nexus-core/internal/config"
)

func Test_buildCommand(t *testing.T) {
	var (
		sub  SubCommand = "bar"
		args            = []string{"--arg1", "--arg2"}
	)

	c := buildCommand("/dummy/dir", "./foo", sub, args, context.TODO())

	if c.Cmd.Dir != "/dummy/dir" {
		t.Fatalf("wrong working directory: %s", c.Cmd.Path)
	}

	if c.Cmd.String() != "./foo bar --arg1 --arg2 --json" {
		t.Fatalf("wrong command: %s", c.Cmd.String())
	}
}

func Test_GetCommandInit(t *testing.T) {
	bp := BaseProvisioner{
		ExecutablePath:   "/usr/local/bin/terraform",
		WorkingDirectory: "bar",
		ProvisionerConfig: config.Provisioner{
			AllowedExecutables: "/usr/local/bin/terraform",
		},
	}

	c, err := bp.GetCommandInit(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if c.Cmd.String() != "/usr/local/bin/terraform init --json" {
		t.Fatalf("wrong command: %s", c.Cmd.String())
	}
}

func Test_GetCommandPlan(t *testing.T) {
	bp := BaseProvisioner{
		ExecutablePath:   "/usr/local/bin/terraform",
		WorkingDirectory: "bar",
		ProvisionerConfig: config.Provisioner{
			AllowedExecutables: "/usr/local/bin/terraform",
		},
	}

	c, err := bp.GetCommandPlan(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if c.Cmd.String() != "/usr/local/bin/terraform plan --json" {
		t.Fatalf("wrong command: %s", c.Cmd.String())
	}
}

func Test_GetCommandApply(t *testing.T) {
	bp := BaseProvisioner{
		ExecutablePath:   "/usr/local/bin/terraform",
		WorkingDirectory: "bar",
		ProvisionerConfig: config.Provisioner{
			AllowedExecutables: "/usr/local/bin/terraform",
		},
	}

	c, err := bp.GetCommandApply(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if c.Cmd.String() != "/usr/local/bin/terraform apply --json" {
		t.Fatalf("wrong command: %s", c.Cmd.String())
	}
}
