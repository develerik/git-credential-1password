package onepassword

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// StoreCredentials saves new credentials to 1password.
func (c *Client) StoreCredentials(protocol, host, username, password string) error {
	creds, err := c.GetCredentials(host)

	if err != nil || creds != nil {
		return nil
	}

	var stdout bytes.Buffer

	var stderr bytes.Buffer

	// TODO: handle session expired error
	l := login{
		Notes: fmt.Sprintf("Protocol: %s", protocol),
		Fields: []field{
			{
				Name:        "username",
				Value:       username,
				Type:        "T",
				Designation: "username",
			},
			{
				Name:        "password",
				Value:       password,
				Type:        "P",
				Designation: "password",
			},
		},
	}

	data, err := json.Marshal(l)

	if err != nil {
		return err
	}

	// base64url encode data
	encData := base64.StdEncoding.EncodeToString(data)
	encData = strings.ReplaceAll(encData, "+", "-") // 62nd char of encoding
	encData = strings.ReplaceAll(encData, "/", "_") // 63rd char of encoding
	encData = strings.ReplaceAll(encData, "=", "")  // Remove any trailing '='s

	cmd := exec.Command("op", "--session", c.token, // nolint:gosec // TODO: validate
		"create", "item", "Login", encData,
		"--title", host, "--tags", "git")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String()) // nolint:goerr113 // TODO: refactor
	}

	return nil
}
