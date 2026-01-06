package provisioning

import (
	"testing"

	"github.com/tbauriedel/resource-nexus-core/internal/config"
)

func Test_Validate(t *testing.T) {
	bp := BaseProvisioner{
		ExecutablePath:   "terraform",
		WorkingDirectory: "/something/else",
		ProvisionerConfig: config.Provisioner{
			AllowedExecutables: "terraform,tofu",
		},
	}

	// Exec path is valid?
	if err := bp.Validate(); err != nil {
		t.Fatalf("validation failed: %s", err.Error())
	}
}
