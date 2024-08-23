// Package systemd provides the common systemd helper functions.
package systemd

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func SystemdExists() bool {
	p, err := exec.LookPath("systemd")
	if err != nil {
		return false
	}
	return p != ""
}

func SystemctlExists() bool {
	p, err := exec.LookPath("systemctl")
	if err != nil {
		return false
	}
	return p != ""
}

func DaemonReload(ctx context.Context) ([]byte, error) {
	cmdPath, err := exec.LookPath("systemctl")
	if err != nil {
		return nil, err
	}
	cmd := exec.CommandContext(ctx, cmdPath, "daemon-reload")
	return cmd.Output()
}

// CheckVersion returns the systemd version by running `systemd --version`.
func CheckVersion() (string, []string, error) {
	p, err := exec.LookPath("systemd")
	if err != nil {
		return "", nil, fmt.Errorf("systemd version check requires systemd (%w)", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	b, err := exec.CommandContext(ctx, p, "--version").CombinedOutput()
	cancel()
	if err != nil {
		return "", nil, err
	}

	ver, extra := parseVersion(string(b))
	return ver, extra, nil
}

func parseVersion(version string) (string, []string) {
	if len(version) == 0 {
		return "", nil
	}

	ver := ""
	extra := make([]string, 0)
	for _, line := range strings.Split(version, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if ver == "" {
			ver = line
			continue
		}

		extra = append(extra, line)
	}
	return ver, extra
}

// IsActive returns true if the systemd service is active.
func IsActive(service string) (bool, error) {
	p, err := exec.LookPath("systemctl")
	if err != nil {
		return false, fmt.Errorf("systemd active check requires systemctl (%w)", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	b, err := exec.CommandContext(ctx, p, "is-active", service).CombinedOutput()
	cancel()
	if err != nil {
		// e.g., "inactive" with exit status 3
		if strings.Contains(string(b), "inactive") {
			return false, nil
		}
		return false, err
	}
	return strings.TrimSpace(string(b)) == "active", nil
}

const uptimeTimeLayout = "Mon 2006-01-02 15:04:05 MST"

// GetUptime returns the uptime duration of the systemd service.
// ref. https://github.com/kubernetes/node-problem-detector/blob/c4e5400ed6d7ca30d3a803248ae5b55c53557e59/pkg/healthchecker/health_checker_linux.go
func GetUptime(service string) (time.Duration, error) {
	p, err := exec.LookPath("systemctl")
	if err != nil {
		return 0, fmt.Errorf("systemd uptime check requires systemctl (%w)", err)
	}

	// below is copied from https://github.com/kubernetes/node-problem-detector/blob/c4e5400ed6d7ca30d3a803248ae5b55c53557e59/pkg/healthchecker/health_checker_linux.go
	//
	// transition from inactive -> activating and the timestamp is captured.
	// Source : https://www.freedesktop.org/wiki/Software/systemd/dbus/
	// Using ActiveEnterTimestamp resulted in race condition where the service was repeatedly killed by plugin when
	// RestartSec of systemd and invoke interval of plugin got in sync. The service was repeatedly killed in
	// activating state and hence ActiveEnterTimestamp was never updated.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	b, err := exec.CommandContext(ctx, p, "show", service, "--property=InactiveExitTimestamp").CombinedOutput()
	cancel()
	if err != nil {
		return 0, err
	}
	val := strings.Split(string(b), "=")
	if len(val) < 2 {
		return time.Duration(0), errors.New("could not parse the service uptime time correctly")
	}
	return parseSystemdUnitUptime(val[1])
}

func parseSystemdUnitUptime(s string) (time.Duration, error) {
	// e.g., "Wed 2024-02-28 01:29:39 UTC\x0a": extra text: "\x0a" will fail without trim
	t, err := time.Parse(uptimeTimeLayout, strings.Trim(s, "\x0a"))
	if err != nil {
		return time.Duration(0), err
	}
	return time.Since(t), nil
}
