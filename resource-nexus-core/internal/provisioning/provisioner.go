package provisioning

import (
	"fmt"
	"slices"
	"strings"

	"github.com/tbauriedel/resource-nexus-core/internal/config"
)

type Provisioner interface {
	GetProvisionerName() string
}

// BaseProvisioner holds the general utilities for all provisioners.
// It is not meant to be used directly.
type BaseProvisioner struct {
	ProvisionerConfig config.Provisioner
	ExecutablePath    string
	WorkingDirectory  string
}

// Validate takes the defined provisioner settings and validates them.
// Returns an error if the executable path is not allowed.
//
// allowedExecutables is a list of allowed executable paths.
func (bp *BaseProvisioner) Validate() error {
	if !slices.Contains(strings.Split(bp.ProvisionerConfig.AllowedExecutables, ","), bp.ExecutablePath) {
		return fmt.Errorf("invalid executable path provided: %s. contact your administrator for help", bp.ExecutablePath)
	}

	return nil
}
