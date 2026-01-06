package provisioner

import "testing"

func Test_GetProvisionerNameTerraform(t *testing.T) {
	ot := Terraform{}

	if ot.GetProvisionerName() != "terraform" {
		t.Fatalf("provisioner name: %s != terraform", ot.GetProvisionerName())
	}
}
