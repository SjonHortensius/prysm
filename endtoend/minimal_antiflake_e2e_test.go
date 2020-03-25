package endtoend

import (
	"testing"

	ev "github.com/prysmaticlabs/prysm/endtoend/evaluators"
	e2eParams "github.com/prysmaticlabs/prysm/endtoend/params"
	"github.com/prysmaticlabs/prysm/endtoend/types"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/prysmaticlabs/prysm/shared/testutil"
)

func TestEndToEnd_AntiFlake_MinimalConfig(t *testing.T) {
	testutil.ResetCache()
	params.UseMinimalConfig()

	minimalConfig := &types.E2EConfig{
		BeaconFlags:    []string{"--minimal-config", "--custom-genesis-delay=10"},
		ValidatorFlags: []string{"--minimal-config"},
		EpochsToRun:    3,
		TestSync:       false,
		TestSlasher:    false,
		Evaluators: []types.Evaluator{
			ev.PeersConnect,
			ev.ValidatorsAreActive,
		},
	}
	if err := e2eParams.Init(4); err != nil {
		t.Fatal(err)
	}

	// Running this test twice to test the quickest conditions (3 epochs) twice.
	runEndToEndTest(t, minimalConfig)
	runEndToEndTest(t, minimalConfig)
}