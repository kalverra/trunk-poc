package fail

import (
	"os"
	"testing"
)

func TestFail(t *testing.T) {
	if os.Getenv("RUN_QUARANTINED_TESTS") != "true" {
		t.Skip("Flaky test quarantined. Ticket <Jira ticket>. Done automatically by branch-out (https://github.com/smartcontractkit/branch-out)")
	} else {
		t.Logf("'RUN_QUARANTINED_TESTS' set to '%s', running quarantined test", os.Getenv("RUN_QUARANTINED_TESTS"))
	}
	t.Parallel()

	t.Log("I fail")
	t.FailNow()
}
