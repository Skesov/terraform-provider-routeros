package routeros

import "testing"

// TestProviderInternalValidate checks every resource and data source schema
// with the SDK's own validation, without a router.
func TestProviderInternalValidate(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}
