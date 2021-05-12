package upgrade_test

import (
	"testing"

	"github.com/jenkins-x/jx-cli/pkg/cmd/upgrade"
	"github.com/stretchr/testify/require"
)

func TestUpgrade(t *testing.T) {
	t.Parallel()

	_, o := upgrade.NewCmdUpgrade()

	err := o.Run()
	require.NoError(t, err, "failed to run upgrade command")
}
