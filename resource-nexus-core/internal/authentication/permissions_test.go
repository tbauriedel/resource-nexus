package authentication

import "testing"

func TestBuildPermissionString(t *testing.T) {
	actual := BuildPermissionString("security", "user", "create")
	expected := "security:user:create"

	if actual != expected {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
