// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: mds_client.go

package mock

import (
	sync "sync"

	github_com_confluentinc_cli_internal_pkg_config_v3 "github.com/confluentinc/cli/internal/pkg/config/v3"
	github_com_confluentinc_cli_internal_pkg_log "github.com/confluentinc/cli/internal/pkg/log"
	github_com_confluentinc_mds_sdk_go "github.com/confluentinc/mds-sdk-go"
)

// MockMDSClientManager is a mock of MDSClientManager interface
type MockMDSClientManager struct {
	lockGetMDSClient sync.Mutex
	GetMDSClientFunc func(ctx *github_com_confluentinc_cli_internal_pkg_config_v3.Context, caCertPath string, flagChanged bool, url string, logger *github_com_confluentinc_cli_internal_pkg_log.Logger) (*github_com_confluentinc_mds_sdk_go.APIClient, error)

	calls struct {
		GetMDSClient []struct {
			Ctx         *github_com_confluentinc_cli_internal_pkg_config_v3.Context
			CaCertPath  string
			FlagChanged bool
			Url         string
			Logger      *github_com_confluentinc_cli_internal_pkg_log.Logger
		}
	}
}

// GetMDSClient mocks base method by wrapping the associated func.
func (m *MockMDSClientManager) GetMDSClient(ctx *github_com_confluentinc_cli_internal_pkg_config_v3.Context, caCertPath string, flagChanged bool, url string, logger *github_com_confluentinc_cli_internal_pkg_log.Logger) (*github_com_confluentinc_mds_sdk_go.APIClient, error) {
	m.lockGetMDSClient.Lock()
	defer m.lockGetMDSClient.Unlock()

	if m.GetMDSClientFunc == nil {
		panic("mocker: MockMDSClientManager.GetMDSClientFunc is nil but MockMDSClientManager.GetMDSClient was called.")
	}

	call := struct {
		Ctx         *github_com_confluentinc_cli_internal_pkg_config_v3.Context
		CaCertPath  string
		FlagChanged bool
		Url         string
		Logger      *github_com_confluentinc_cli_internal_pkg_log.Logger
	}{
		Ctx:         ctx,
		CaCertPath:  caCertPath,
		FlagChanged: flagChanged,
		Url:         url,
		Logger:      logger,
	}

	m.calls.GetMDSClient = append(m.calls.GetMDSClient, call)

	return m.GetMDSClientFunc(ctx, caCertPath, flagChanged, url, logger)
}

// GetMDSClientCalled returns true if GetMDSClient was called at least once.
func (m *MockMDSClientManager) GetMDSClientCalled() bool {
	m.lockGetMDSClient.Lock()
	defer m.lockGetMDSClient.Unlock()

	return len(m.calls.GetMDSClient) > 0
}

// GetMDSClientCalls returns the calls made to GetMDSClient.
func (m *MockMDSClientManager) GetMDSClientCalls() []struct {
	Ctx         *github_com_confluentinc_cli_internal_pkg_config_v3.Context
	CaCertPath  string
	FlagChanged bool
	Url         string
	Logger      *github_com_confluentinc_cli_internal_pkg_log.Logger
} {
	m.lockGetMDSClient.Lock()
	defer m.lockGetMDSClient.Unlock()

	return m.calls.GetMDSClient
}

// Reset resets the calls made to the mocked methods.
func (m *MockMDSClientManager) Reset() {
	m.lockGetMDSClient.Lock()
	m.calls.GetMDSClient = nil
	m.lockGetMDSClient.Unlock()
}
