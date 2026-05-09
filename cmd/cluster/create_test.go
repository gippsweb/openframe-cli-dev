package cluster

import (
	"testing"

	"github.com/gippsweb/openframe-cli-dev/internal/cluster/utils"
	"github.com/gippsweb/openframe-cli-dev/tests/testutil"
)

func init() {
	testutil.InitializeTestMode()
}

func TestCreateCommand(t *testing.T) {
	setupFunc := func() {
		utils.SetTestExecutor(testutil.NewTestMockExecutor())
	}
	teardownFunc := func() {
		utils.ResetGlobalFlags()
	}

	testutil.TestClusterCommand(t, "create", getCreateCmd, setupFunc, teardownFunc)
}
