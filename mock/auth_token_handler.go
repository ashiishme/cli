// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: internal/pkg/auth/auth_token_handler.go

package mock

import (
	sync "sync"

	github_com_confluentinc_cli_internal_pkg_auth "github.com/confluentinc/cli/v3/pkg/auth"
	github_com_confluentinc_mds_sdk_go_mdsv1 "github.com/confluentinc/mds-sdk-go-public/mdsv1"
)

// AuthTokenHandler is a mock of AuthTokenHandler interface
type AuthTokenHandler struct {
	lockGetCCloudTokens sync.Mutex
	GetCCloudTokensFunc func(clientFactory github_com_confluentinc_cli_internal_pkg_auth.CCloudClientFactory, url string, credentials *github_com_confluentinc_cli_internal_pkg_auth.Credentials, noBrowser bool, orgResourceId string) (string, string, error)

	lockGetConfluentToken sync.Mutex
	GetConfluentTokenFunc func(mdsClient *github_com_confluentinc_mds_sdk_go_mdsv1.APIClient, credentials *github_com_confluentinc_cli_internal_pkg_auth.Credentials) (string, error)

	calls struct {
		GetCCloudTokens []struct {
			ClientFactory github_com_confluentinc_cli_internal_pkg_auth.CCloudClientFactory
			Url           string
			Credentials   *github_com_confluentinc_cli_internal_pkg_auth.Credentials
			NoBrowser     bool
			OrgResourceId string
		}
		GetConfluentToken []struct {
			MdsClient   *github_com_confluentinc_mds_sdk_go_mdsv1.APIClient
			Credentials *github_com_confluentinc_cli_internal_pkg_auth.Credentials
		}
	}
}

// GetCCloudTokens mocks base method by wrapping the associated func.
func (m *AuthTokenHandler) GetCCloudTokens(clientFactory github_com_confluentinc_cli_internal_pkg_auth.CCloudClientFactory, url string, credentials *github_com_confluentinc_cli_internal_pkg_auth.Credentials, noBrowser bool, orgResourceId string) (string, string, error) {
	m.lockGetCCloudTokens.Lock()
	defer m.lockGetCCloudTokens.Unlock()

	if m.GetCCloudTokensFunc == nil {
		panic("mocker: AuthTokenHandler.GetCCloudTokensFunc is nil but AuthTokenHandler.GetCCloudTokens was called.")
	}

	call := struct {
		ClientFactory github_com_confluentinc_cli_internal_pkg_auth.CCloudClientFactory
		Url           string
		Credentials   *github_com_confluentinc_cli_internal_pkg_auth.Credentials
		NoBrowser     bool
		OrgResourceId string
	}{
		ClientFactory: clientFactory,
		Url:           url,
		Credentials:   credentials,
		NoBrowser:     noBrowser,
		OrgResourceId: orgResourceId,
	}

	m.calls.GetCCloudTokens = append(m.calls.GetCCloudTokens, call)

	return m.GetCCloudTokensFunc(clientFactory, url, credentials, noBrowser, orgResourceId)
}

// GetCCloudTokensCalled returns true if GetCCloudTokens was called at least once.
func (m *AuthTokenHandler) GetCCloudTokensCalled() bool {
	m.lockGetCCloudTokens.Lock()
	defer m.lockGetCCloudTokens.Unlock()

	return len(m.calls.GetCCloudTokens) > 0
}

// GetCCloudTokensCalls returns the calls made to GetCCloudTokens.
func (m *AuthTokenHandler) GetCCloudTokensCalls() []struct {
	ClientFactory github_com_confluentinc_cli_internal_pkg_auth.CCloudClientFactory
	Url           string
	Credentials   *github_com_confluentinc_cli_internal_pkg_auth.Credentials
	NoBrowser     bool
	OrgResourceId string
} {
	m.lockGetCCloudTokens.Lock()
	defer m.lockGetCCloudTokens.Unlock()

	return m.calls.GetCCloudTokens
}

// GetConfluentToken mocks base method by wrapping the associated func.
func (m *AuthTokenHandler) GetConfluentToken(mdsClient *github_com_confluentinc_mds_sdk_go_mdsv1.APIClient, credentials *github_com_confluentinc_cli_internal_pkg_auth.Credentials) (string, error) {
	m.lockGetConfluentToken.Lock()
	defer m.lockGetConfluentToken.Unlock()

	if m.GetConfluentTokenFunc == nil {
		panic("mocker: AuthTokenHandler.GetConfluentTokenFunc is nil but AuthTokenHandler.GetConfluentToken was called.")
	}

	call := struct {
		MdsClient   *github_com_confluentinc_mds_sdk_go_mdsv1.APIClient
		Credentials *github_com_confluentinc_cli_internal_pkg_auth.Credentials
	}{
		MdsClient:   mdsClient,
		Credentials: credentials,
	}

	m.calls.GetConfluentToken = append(m.calls.GetConfluentToken, call)

	return m.GetConfluentTokenFunc(mdsClient, credentials)
}

// GetConfluentTokenCalled returns true if GetConfluentToken was called at least once.
func (m *AuthTokenHandler) GetConfluentTokenCalled() bool {
	m.lockGetConfluentToken.Lock()
	defer m.lockGetConfluentToken.Unlock()

	return len(m.calls.GetConfluentToken) > 0
}

// GetConfluentTokenCalls returns the calls made to GetConfluentToken.
func (m *AuthTokenHandler) GetConfluentTokenCalls() []struct {
	MdsClient   *github_com_confluentinc_mds_sdk_go_mdsv1.APIClient
	Credentials *github_com_confluentinc_cli_internal_pkg_auth.Credentials
} {
	m.lockGetConfluentToken.Lock()
	defer m.lockGetConfluentToken.Unlock()

	return m.calls.GetConfluentToken
}

// Reset resets the calls made to the mocked methods.
func (m *AuthTokenHandler) Reset() {
	m.lockGetCCloudTokens.Lock()
	m.calls.GetCCloudTokens = nil
	m.lockGetCCloudTokens.Unlock()
	m.lockGetConfluentToken.Lock()
	m.calls.GetConfluentToken = nil
	m.lockGetConfluentToken.Unlock()
}
