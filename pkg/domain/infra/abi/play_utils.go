package abi

import (
	"fmt"

	"github.com/containers/common/pkg/config"
	"github.com/containers/podman/v4/libpod/define"
)

// getSdNotifyMode returns the `sdNotifyAnnotation/$name` for the specified
// name. If name is empty, it'll only look for `sdNotifyAnnotation`.
func getSdNotifyMode(annotations map[string]string, name string) (string, error) {
	var mode string
	switch len(name) {
	case 0:
		mode = annotations[sdNotifyAnnotation]
	default:
		mode = annotations[sdNotifyAnnotation+"/"+name]
	}
	return mode, define.ValidateSdNotifyMode(mode)
}

// getPodExitPolicy returns the pod's exit policy as set in the annotations.
func getPodExitPolicy(annotations map[string]string) (string, error) {
	var policy string
	policy = annotations[define.PodExitPolicy]
	if policy == "" {
		policy = string(config.PodExitPolicyStop)
	}
	switch policy {
	case string(config.PodExitPolicyContinue), string(config.PodExitPolicyStop), string(config.PodExitPolicyImmediate):
        return policy, nil
    default:
		return policy, fmt.Errorf("%w: invalid exit policy value %q: must be %s, %s or %s", define.ErrInvalidArg, policy, config.PodExitPolicyContinue, config.PodExitPolicyStop, config.PodExitPolicyImmediate)
	}
}
