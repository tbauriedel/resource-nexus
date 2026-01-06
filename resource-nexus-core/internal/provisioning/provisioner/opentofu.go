package provisioner

import "github.com/tbauriedel/resource-nexus-core/internal/provisioning"

type OpenTofu struct {
	provisioning.BaseProvisioner
}

// GetProvisionerName returns the provisioner name
//
// Needed to implement the provisioning.Provisioner interface.
func (ot *OpenTofu) GetProvisionerName() string {
	return "opentofu"
}
