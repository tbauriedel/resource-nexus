package provisioner

import "testing"

func Test_GetProvisionerNameOpenTofu(t *testing.T) {
	ot := OpenTofu{}

	if ot.GetProvisionerName() != "opentofu" {
		t.Fatalf("provisioner name: %s != opentofu", ot.GetProvisionerName())
	}
}
