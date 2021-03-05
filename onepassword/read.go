package onepassword

import (
	"bytes"
	"encoding/json"
	"errors"

	exec "golang.org/x/sys/execabs"
)

// GetCredentials loads credentials from 1password.
func (c *Client) GetCredentials(host string) (*Credentials, error) {
	var stdout bytes.Buffer

	var stderr bytes.Buffer

	// TODO: handle session expired error
	cmd := exec.Command("op", "--cache", "--session", c.token, "get", "item", host) // nolint:gosec // TODO: validate
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, errors.New(stderr.String()) // nolint:goerr113 // TODO: refactor
	}

	var res response
	if err := json.Unmarshal(stdout.Bytes(), &res); err != nil {
		return nil, err
	}

	var username string

	var password string

	for _, field := range res.Details.Fields {
		if field.Name == "username" {
			username = field.Value
		}

		if field.Name == "password" {
			password = field.Value
		}
	}

	return &Credentials{
		Username: username,
		Password: password,
	}, nil
}
