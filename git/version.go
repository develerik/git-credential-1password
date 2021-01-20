package git

import (
	"bytes"
	"errors"
	"strings"

	exec "golang.org/x/sys/execabs"
)

// GetVersion returns the installed git version.
func GetVersion() (string, error) {
	var stdout bytes.Buffer

	var stderr bytes.Buffer

	cmd := exec.Command("git", "--version")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", errors.New(stderr.String()) // nolint:goerr113 // TODO: refactor
	}

	v := strings.Split(stdout.String(), " ")

	if len(v) != 3 { // nolint:gomnd // git version output parts
		return "", errors.New("could not parse git version") // nolint:goerr113 // TODO: refactor
	}

	version := strings.TrimSuffix(v[2], "\n")

	return version, nil
}
