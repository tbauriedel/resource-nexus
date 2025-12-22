package tf

import (
	"fmt"
	"testing"
)

func TestBuildCommand(t *testing.T) {
	tf := getDefaults()
	tf.ExecutablePath = "./terraform"
	tf.tmpWorkDir = "/tmp/foo"

	cmd := tf.buildCommand("dummy", []string{"--foo", "--bar"})

	if cmd.Cmd.String() != "./terraform dummy --foo --bar --json" {
		t.Fatalf("command does not match expected value\nactual: %s\nexpected: ./terraform dummy --foo bar", cmd.Cmd.String())
	}
}

func TestGetCommandInit(t *testing.T) {
	tf := getDefaults()
	tf.ExecutablePath = "./terraform"
	tf.tmpWorkDir = "/tmp/foo"

	c := tf.GetCommandInit()
	if c.Cmd.String() != fmt.Sprintf("./terraform %s --json", SubCommandInit) {
		t.Fatal("command does not match expected value")
	}
}

func TestCommandPlan(t *testing.T) {
	tf := getDefaults()
	tf.ExecutablePath = "./terraform"
	tf.tmpWorkDir = "/tmp/foo"

	c := tf.GetCommandPlan()
	if c.Cmd.String() != fmt.Sprintf("./terraform %s --json", SubCommandPlan) {
		t.Fatal("command does not match expected value")
	}
}

func TestCommandApply(t *testing.T) {
	tf := getDefaults()
	tf.ExecutablePath = "./terraform"
	tf.tmpWorkDir = "/tmp/foo"

	c := tf.GetCommandApply()
	if c.Cmd.String() != fmt.Sprintf("./terraform %s --json", SubCommandApply) {
		t.Fatal("command does not match expected value")
	}
}
