// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: auth_token_handler.go

package mock

import (
	sync "sync"

	github_com_confluentinc_cli_v3_pkg_auth "github.com/confluentinc/cli/v3/pkg/auth"
	github_com_confluentinc_mds_sdk_go_public_mdsv1 "github.com/confluentinc/mds-sdk-go-public/mdsv1"
)

// AuthTokenHandler is a mock of AuthTokenHandler interface
type AuthTokenHandler struct {
	lockGetCCloudTokens sync.Mutex
	GetCCloudTokensFunc func(arg0 github_com_confluentinc_cli_v3_pkg_auth.CCloudClientFactory, arg1 string, arg2 *github_com_confluentinc_cli_v3_pkg_auth.Credentials, arg3 bool, arg4 string) (string, string, error)

	lockGetConfluentToken sync.Mutex
	GetConfluentTokenFunc func(arg0 *github_com_confluentinc_mds_sdk_go_public_mdsv1.APIClient, arg1 *github_com_confluentinc_cli_v3_pkg_auth.Credentials, arg2 bool) (string, string, error)

	calls struct {
		GetCCloudTokens []struct {
			Arg0 github_com_confluentinc_cli_v3_pkg_auth.CCloudClientFactory
			Arg1 string
			Arg2 *github_com_confluentinc_cli_v3_pkg_auth.Credentials
			Arg3 bool
			Arg4 string
		}
		GetConfluentToken []struct {
			Arg0 *github_com_confluentinc_mds_sdk_go_public_mdsv1.APIClient
			Arg1 *github_com_confluentinc_cli_v3_pkg_auth.Credentials
			Arg2 bool
		}
	}
}

// GetCCloudTokens mocks base method by wrapping the associated func.
func (m *AuthTokenHandler) GetCCloudTokens(arg0 github_com_confluentinc_cli_v3_pkg_auth.CCloudClientFactory, arg1 string, arg2 *github_com_confluentinc_cli_v3_pkg_auth.Credentials, arg3 bool, arg4 string) (string, string, error) {
	m.lockGetCCloudTokens.Lock()
	defer m.lockGetCCloudTokens.Unlock()

	if m.GetCCloudTokensFunc == nil {
		panic("mocker: AuthTokenHandler.GetCCloudTokensFunc is nil but AuthTokenHandler.GetCCloudTokens was called.")
	}

	call := struct {
		Arg0 github_com_confluentinc_cli_v3_pkg_auth.CCloudClientFactory
		Arg1 string
		Arg2 *github_com_confluentinc_cli_v3_pkg_auth.Credentials
		Arg3 bool
		Arg4 string
	}{
		Arg0: arg0,
		Arg1: arg1,
		Arg2: arg2,
		Arg3: arg3,
		Arg4: arg4,
	}

	m.calls.GetCCloudTokens = append(m.calls.GetCCloudTokens, call)

	return m.GetCCloudTokensFunc(arg0, arg1, arg2, arg3, arg4)
}

// GetCCloudTokensCalled returns true if GetCCloudTokens was called at least once.
func (m *AuthTokenHandler) GetCCloudTokensCalled() bool {
	m.lockGetCCloudTokens.Lock()
	defer m.lockGetCCloudTokens.Unlock()

	return len(m.calls.GetCCloudTokens) > 0
}

// GetCCloudTokensCalls returns the calls made to GetCCloudTokens.
func (m *AuthTokenHandler) GetCCloudTokensCalls() []struct {
	Arg0 github_com_confluentinc_cli_v3_pkg_auth.CCloudClientFactory
	Arg1 string
	Arg2 *github_com_confluentinc_cli_v3_pkg_auth.Credentials
	Arg3 bool
	Arg4 string
} {
	m.lockGetCCloudTokens.Lock()
	defer m.lockGetCCloudTokens.Unlock()

	return m.calls.GetCCloudTokens
}

// GetConfluentToken mocks base method by wrapping the associated func.
func (m *AuthTokenHandler) GetConfluentToken(arg0 *github_com_confluentinc_mds_sdk_go_public_mdsv1.APIClient, arg1 *github_com_confluentinc_cli_v3_pkg_auth.Credentials, arg2 bool) (string, string, error) {
	m.lockGetConfluentToken.Lock()
	defer m.lockGetConfluentToken.Unlock()

	if m.GetConfluentTokenFunc == nil {
		panic("mocker: AuthTokenHandler.GetConfluentTokenFunc is nil but AuthTokenHandler.GetConfluentToken was called.")
	}

	call := struct {
		Arg0 *github_com_confluentinc_mds_sdk_go_public_mdsv1.APIClient
		Arg1 *github_com_confluentinc_cli_v3_pkg_auth.Credentials
		Arg2 bool
	}{
		Arg0: arg0,
		Arg1: arg1,
		Arg2: arg2,
	}

	m.calls.GetConfluentToken = append(m.calls.GetConfluentToken, call)

	return m.GetConfluentTokenFunc(arg0, arg1, arg2)
}

// GetConfluentTokenCalled returns true if GetConfluentToken was called at least once.
func (m *AuthTokenHandler) GetConfluentTokenCalled() bool {
	m.lockGetConfluentToken.Lock()
	defer m.lockGetConfluentToken.Unlock()

	return len(m.calls.GetConfluentToken) > 0
}

// GetConfluentTokenCalls returns the calls made to GetConfluentToken.
func (m *AuthTokenHandler) GetConfluentTokenCalls() []struct {
	Arg0 *github_com_confluentinc_mds_sdk_go_public_mdsv1.APIClient
	Arg1 *github_com_confluentinc_cli_v3_pkg_auth.Credentials
	Arg2 bool
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
