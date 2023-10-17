// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: client.go

package mock

import (
	sync "sync"
)

// Client is a mock of Client interface
type Client struct {
	lockCheckForUpdates sync.Mutex
	CheckForUpdatesFunc func(cliName, currentVersion string, forceCheck bool) (string, string, error)

	lockGetLatestReleaseNotes sync.Mutex
	GetLatestReleaseNotesFunc func(cliName, currentVersion string) (string, []string, error)

	lockPromptToDownload sync.Mutex
	PromptToDownloadFunc func(cliName, currVersion, latestVersion, releaseNotes string, confirm bool) bool

	lockUpdateBinary sync.Mutex
	UpdateBinaryFunc func(cliName, version string, noVerify bool) error

	calls struct {
		CheckForUpdates []struct {
			CliName        string
			CurrentVersion string
			ForceCheck     bool
		}
		GetLatestReleaseNotes []struct {
			CliName        string
			CurrentVersion string
		}
		PromptToDownload []struct {
			CliName       string
			CurrVersion   string
			LatestVersion string
			ReleaseNotes  string
			Confirm       bool
		}
		UpdateBinary []struct {
			CliName  string
			Version  string
			NoVerify bool
		}
	}
}

// CheckForUpdates mocks base method by wrapping the associated func.
func (m *Client) CheckForUpdates(cliName, currentVersion string, forceCheck bool) (string, string, error) {
	m.lockCheckForUpdates.Lock()
	defer m.lockCheckForUpdates.Unlock()

	if m.CheckForUpdatesFunc == nil {
		panic("mocker: Client.CheckForUpdatesFunc is nil but Client.CheckForUpdates was called.")
	}

	call := struct {
		CliName        string
		CurrentVersion string
		ForceCheck     bool
	}{
		CliName:        cliName,
		CurrentVersion: currentVersion,
		ForceCheck:     forceCheck,
	}

	m.calls.CheckForUpdates = append(m.calls.CheckForUpdates, call)

	return m.CheckForUpdatesFunc(cliName, currentVersion, forceCheck)
}

// CheckForUpdatesCalled returns true if CheckForUpdates was called at least once.
func (m *Client) CheckForUpdatesCalled() bool {
	m.lockCheckForUpdates.Lock()
	defer m.lockCheckForUpdates.Unlock()

	return len(m.calls.CheckForUpdates) > 0
}

// CheckForUpdatesCalls returns the calls made to CheckForUpdates.
func (m *Client) CheckForUpdatesCalls() []struct {
	CliName        string
	CurrentVersion string
	ForceCheck     bool
} {
	m.lockCheckForUpdates.Lock()
	defer m.lockCheckForUpdates.Unlock()

	return m.calls.CheckForUpdates
}

// GetLatestReleaseNotes mocks base method by wrapping the associated func.
func (m *Client) GetLatestReleaseNotes(cliName, currentVersion string) (string, []string, error) {
	m.lockGetLatestReleaseNotes.Lock()
	defer m.lockGetLatestReleaseNotes.Unlock()

	if m.GetLatestReleaseNotesFunc == nil {
		panic("mocker: Client.GetLatestReleaseNotesFunc is nil but Client.GetLatestReleaseNotes was called.")
	}

	call := struct {
		CliName        string
		CurrentVersion string
	}{
		CliName:        cliName,
		CurrentVersion: currentVersion,
	}

	m.calls.GetLatestReleaseNotes = append(m.calls.GetLatestReleaseNotes, call)

	return m.GetLatestReleaseNotesFunc(cliName, currentVersion)
}

// GetLatestReleaseNotesCalled returns true if GetLatestReleaseNotes was called at least once.
func (m *Client) GetLatestReleaseNotesCalled() bool {
	m.lockGetLatestReleaseNotes.Lock()
	defer m.lockGetLatestReleaseNotes.Unlock()

	return len(m.calls.GetLatestReleaseNotes) > 0
}

// GetLatestReleaseNotesCalls returns the calls made to GetLatestReleaseNotes.
func (m *Client) GetLatestReleaseNotesCalls() []struct {
	CliName        string
	CurrentVersion string
} {
	m.lockGetLatestReleaseNotes.Lock()
	defer m.lockGetLatestReleaseNotes.Unlock()

	return m.calls.GetLatestReleaseNotes
}

// PromptToDownload mocks base method by wrapping the associated func.
func (m *Client) PromptToDownload(cliName, currVersion, latestVersion, releaseNotes string, confirm bool) bool {
	m.lockPromptToDownload.Lock()
	defer m.lockPromptToDownload.Unlock()

	if m.PromptToDownloadFunc == nil {
		panic("mocker: Client.PromptToDownloadFunc is nil but Client.PromptToDownload was called.")
	}

	call := struct {
		CliName       string
		CurrVersion   string
		LatestVersion string
		ReleaseNotes  string
		Confirm       bool
	}{
		CliName:       cliName,
		CurrVersion:   currVersion,
		LatestVersion: latestVersion,
		ReleaseNotes:  releaseNotes,
		Confirm:       confirm,
	}

	m.calls.PromptToDownload = append(m.calls.PromptToDownload, call)

	return m.PromptToDownloadFunc(cliName, currVersion, latestVersion, releaseNotes, confirm)
}

// PromptToDownloadCalled returns true if PromptToDownload was called at least once.
func (m *Client) PromptToDownloadCalled() bool {
	m.lockPromptToDownload.Lock()
	defer m.lockPromptToDownload.Unlock()

	return len(m.calls.PromptToDownload) > 0
}

// PromptToDownloadCalls returns the calls made to PromptToDownload.
func (m *Client) PromptToDownloadCalls() []struct {
	CliName       string
	CurrVersion   string
	LatestVersion string
	ReleaseNotes  string
	Confirm       bool
} {
	m.lockPromptToDownload.Lock()
	defer m.lockPromptToDownload.Unlock()

	return m.calls.PromptToDownload
}

// UpdateBinary mocks base method by wrapping the associated func.
func (m *Client) UpdateBinary(cliName, version string, noVerify bool) error {
	m.lockUpdateBinary.Lock()
	defer m.lockUpdateBinary.Unlock()

	if m.UpdateBinaryFunc == nil {
		panic("mocker: Client.UpdateBinaryFunc is nil but Client.UpdateBinary was called.")
	}

	call := struct {
		CliName  string
		Version  string
		NoVerify bool
	}{
		CliName:  cliName,
		Version:  version,
		NoVerify: noVerify,
	}

	m.calls.UpdateBinary = append(m.calls.UpdateBinary, call)

	return m.UpdateBinaryFunc(cliName, version, noVerify)
}

// UpdateBinaryCalled returns true if UpdateBinary was called at least once.
func (m *Client) UpdateBinaryCalled() bool {
	m.lockUpdateBinary.Lock()
	defer m.lockUpdateBinary.Unlock()

	return len(m.calls.UpdateBinary) > 0
}

// UpdateBinaryCalls returns the calls made to UpdateBinary.
func (m *Client) UpdateBinaryCalls() []struct {
	CliName  string
	Version  string
	NoVerify bool
} {
	m.lockUpdateBinary.Lock()
	defer m.lockUpdateBinary.Unlock()

	return m.calls.UpdateBinary
}

// Reset resets the calls made to the mocked methods.
func (m *Client) Reset() {
	m.lockCheckForUpdates.Lock()
	m.calls.CheckForUpdates = nil
	m.lockCheckForUpdates.Unlock()
	m.lockGetLatestReleaseNotes.Lock()
	m.calls.GetLatestReleaseNotes = nil
	m.lockGetLatestReleaseNotes.Unlock()
	m.lockPromptToDownload.Lock()
	m.calls.PromptToDownload = nil
	m.lockPromptToDownload.Unlock()
	m.lockUpdateBinary.Lock()
	m.calls.UpdateBinary = nil
	m.lockUpdateBinary.Unlock()
}
