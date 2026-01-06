package provisioner

import "github.com/tbauriedel/resource-nexus-core/internal/provisioning"

type Terraform struct {
	provisioning.BaseProvisioner
}

// GetProvisionerName returns the provisioner name
//
// Needed to implement the provisioning.Provisioner interface.
func (tf *Terraform) GetProvisionerName() string {
	return "terraform"
}
