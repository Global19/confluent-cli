// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: update_token_handler.go

package mock

import (
	sync "sync"

	github_com_confluentinc_cli_internal_pkg_config_v3 "github.com/confluentinc/cli/internal/pkg/config/v3"
	github_com_confluentinc_cli_internal_pkg_log "github.com/confluentinc/cli/internal/pkg/log"
)

// MockUpdateTokenHandler is a mock of UpdateTokenHandler interface
type MockUpdateTokenHandler struct {
	lockUpdateCCloudAuthTokenUsingNetrcCredentials sync.Mutex
	UpdateCCloudAuthTokenUsingNetrcCredentialsFunc func(ctx *github_com_confluentinc_cli_internal_pkg_config_v3.Context, userAgent string, logger *github_com_confluentinc_cli_internal_pkg_log.Logger) error

	lockUpdateConfluentAuthTokenUsingNetrcCredentials sync.Mutex
	UpdateConfluentAuthTokenUsingNetrcCredentialsFunc func(ctx *github_com_confluentinc_cli_internal_pkg_config_v3.Context, logger *github_com_confluentinc_cli_internal_pkg_log.Logger) error

	calls struct {
		UpdateCCloudAuthTokenUsingNetrcCredentials []struct {
			Ctx       *github_com_confluentinc_cli_internal_pkg_config_v3.Context
			UserAgent string
			Logger    *github_com_confluentinc_cli_internal_pkg_log.Logger
		}
		UpdateConfluentAuthTokenUsingNetrcCredentials []struct {
			Ctx    *github_com_confluentinc_cli_internal_pkg_config_v3.Context
			Logger *github_com_confluentinc_cli_internal_pkg_log.Logger
		}
	}
}

// UpdateCCloudAuthTokenUsingNetrcCredentials mocks base method by wrapping the associated func.
func (m *MockUpdateTokenHandler) UpdateCCloudAuthTokenUsingNetrcCredentials(ctx *github_com_confluentinc_cli_internal_pkg_config_v3.Context, userAgent string, logger *github_com_confluentinc_cli_internal_pkg_log.Logger) error {
	m.lockUpdateCCloudAuthTokenUsingNetrcCredentials.Lock()
	defer m.lockUpdateCCloudAuthTokenUsingNetrcCredentials.Unlock()

	if m.UpdateCCloudAuthTokenUsingNetrcCredentialsFunc == nil {
		panic("mocker: MockUpdateTokenHandler.UpdateCCloudAuthTokenUsingNetrcCredentialsFunc is nil but MockUpdateTokenHandler.UpdateCCloudAuthTokenUsingNetrcCredentials was called.")
	}

	call := struct {
		Ctx       *github_com_confluentinc_cli_internal_pkg_config_v3.Context
		UserAgent string
		Logger    *github_com_confluentinc_cli_internal_pkg_log.Logger
	}{
		Ctx:       ctx,
		UserAgent: userAgent,
		Logger:    logger,
	}

	m.calls.UpdateCCloudAuthTokenUsingNetrcCredentials = append(m.calls.UpdateCCloudAuthTokenUsingNetrcCredentials, call)

	return m.UpdateCCloudAuthTokenUsingNetrcCredentialsFunc(ctx, userAgent, logger)
}

// UpdateCCloudAuthTokenUsingNetrcCredentialsCalled returns true if UpdateCCloudAuthTokenUsingNetrcCredentials was called at least once.
func (m *MockUpdateTokenHandler) UpdateCCloudAuthTokenUsingNetrcCredentialsCalled() bool {
	m.lockUpdateCCloudAuthTokenUsingNetrcCredentials.Lock()
	defer m.lockUpdateCCloudAuthTokenUsingNetrcCredentials.Unlock()

	return len(m.calls.UpdateCCloudAuthTokenUsingNetrcCredentials) > 0
}

// UpdateCCloudAuthTokenUsingNetrcCredentialsCalls returns the calls made to UpdateCCloudAuthTokenUsingNetrcCredentials.
func (m *MockUpdateTokenHandler) UpdateCCloudAuthTokenUsingNetrcCredentialsCalls() []struct {
	Ctx       *github_com_confluentinc_cli_internal_pkg_config_v3.Context
	UserAgent string
	Logger    *github_com_confluentinc_cli_internal_pkg_log.Logger
} {
	m.lockUpdateCCloudAuthTokenUsingNetrcCredentials.Lock()
	defer m.lockUpdateCCloudAuthTokenUsingNetrcCredentials.Unlock()

	return m.calls.UpdateCCloudAuthTokenUsingNetrcCredentials
}

// UpdateConfluentAuthTokenUsingNetrcCredentials mocks base method by wrapping the associated func.
func (m *MockUpdateTokenHandler) UpdateConfluentAuthTokenUsingNetrcCredentials(ctx *github_com_confluentinc_cli_internal_pkg_config_v3.Context, logger *github_com_confluentinc_cli_internal_pkg_log.Logger) error {
	m.lockUpdateConfluentAuthTokenUsingNetrcCredentials.Lock()
	defer m.lockUpdateConfluentAuthTokenUsingNetrcCredentials.Unlock()

	if m.UpdateConfluentAuthTokenUsingNetrcCredentialsFunc == nil {
		panic("mocker: MockUpdateTokenHandler.UpdateConfluentAuthTokenUsingNetrcCredentialsFunc is nil but MockUpdateTokenHandler.UpdateConfluentAuthTokenUsingNetrcCredentials was called.")
	}

	call := struct {
		Ctx    *github_com_confluentinc_cli_internal_pkg_config_v3.Context
		Logger *github_com_confluentinc_cli_internal_pkg_log.Logger
	}{
		Ctx:    ctx,
		Logger: logger,
	}

	m.calls.UpdateConfluentAuthTokenUsingNetrcCredentials = append(m.calls.UpdateConfluentAuthTokenUsingNetrcCredentials, call)

	return m.UpdateConfluentAuthTokenUsingNetrcCredentialsFunc(ctx, logger)
}

// UpdateConfluentAuthTokenUsingNetrcCredentialsCalled returns true if UpdateConfluentAuthTokenUsingNetrcCredentials was called at least once.
func (m *MockUpdateTokenHandler) UpdateConfluentAuthTokenUsingNetrcCredentialsCalled() bool {
	m.lockUpdateConfluentAuthTokenUsingNetrcCredentials.Lock()
	defer m.lockUpdateConfluentAuthTokenUsingNetrcCredentials.Unlock()

	return len(m.calls.UpdateConfluentAuthTokenUsingNetrcCredentials) > 0
}

// UpdateConfluentAuthTokenUsingNetrcCredentialsCalls returns the calls made to UpdateConfluentAuthTokenUsingNetrcCredentials.
func (m *MockUpdateTokenHandler) UpdateConfluentAuthTokenUsingNetrcCredentialsCalls() []struct {
	Ctx    *github_com_confluentinc_cli_internal_pkg_config_v3.Context
	Logger *github_com_confluentinc_cli_internal_pkg_log.Logger
} {
	m.lockUpdateConfluentAuthTokenUsingNetrcCredentials.Lock()
	defer m.lockUpdateConfluentAuthTokenUsingNetrcCredentials.Unlock()

	return m.calls.UpdateConfluentAuthTokenUsingNetrcCredentials
}

// Reset resets the calls made to the mocked methods.
func (m *MockUpdateTokenHandler) Reset() {
	m.lockUpdateCCloudAuthTokenUsingNetrcCredentials.Lock()
	m.calls.UpdateCCloudAuthTokenUsingNetrcCredentials = nil
	m.lockUpdateCCloudAuthTokenUsingNetrcCredentials.Unlock()
	m.lockUpdateConfluentAuthTokenUsingNetrcCredentials.Lock()
	m.calls.UpdateConfluentAuthTokenUsingNetrcCredentials = nil
	m.lockUpdateConfluentAuthTokenUsingNetrcCredentials.Unlock()
}
