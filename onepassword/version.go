package onepassword

import (
	"bytes"
	"errors"
	"strings"

	exec "golang.org/x/sys/execabs"
)

// GetVersion returns the installed 1password cli version.
func GetVersion() (string, error) {
	var stdout bytes.Buffer

	var stderr bytes.Buffer

	cmd := exec.Command("op", "--version")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", errors.New(stderr.String()) // nolint:goerr113 // TODO: refactor
	}

	version := strings.TrimSuffix(stdout.String(), "\n")

	return version, nil
}
