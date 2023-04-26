package google

import (
	"fmt"
	"time"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

type NetworkSecurityAddressGroupOperationWaiter struct {
	Config    *transport_tpg.Config
	UserAgent string
	CommonOperationWaiter
}

func (w *NetworkSecurityAddressGroupOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("%s%s", w.Config.NetworkSecurityBasePath, w.CommonOperationWaiter.Op.Name)

	return SendRequest(w.Config, "GET", "", url, w.UserAgent, nil)
}

func createNetworkSecurityAddressGroupWaiter(config *transport_tpg.Config, op map[string]interface{}, activity, userAgent string) (*NetworkSecurityAddressGroupOperationWaiter, error) {
	w := &NetworkSecurityAddressGroupOperationWaiter{
		Config:    config,
		UserAgent: userAgent,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

func NetworkSecurityAddressGroupOperationWaitTime(config *transport_tpg.Config, op map[string]interface{}, activity, userAgent string, timeout time.Duration) error {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil
	}
	w, err := createNetworkSecurityAddressGroupWaiter(config, op, activity, userAgent)
	if err != nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
