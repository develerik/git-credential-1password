package onepassword

import (
	"bytes"
	"errors"
	"os/exec"
)

// DeleteCredentials deletes credentials from 1password.
func (c *Client) DeleteCredentials(_, host string) error {
	var stdout bytes.Buffer

	var stderr bytes.Buffer

	cmd := exec.Command("op", "--session", c.token, // nolint:gosec // TODO: validate
		"delete", "item", host)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String()) // nolint:goerr113 // TODO: refactor
	}

	return nil
}
