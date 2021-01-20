package git

import (
	"bytes"
	"errors"
	"fmt"

	exec "golang.org/x/sys/execabs"
)

var eof []byte

// GetFromCache reads a 1password session token from git cache.
func GetFromCache(account string) (string, error) {
	var stdout bytes.Buffer

	var stdin bytes.Buffer

	var stderr bytes.Buffer

	cmd := exec.Command("git", "credential-cache", "get")
	cmd.Stdout = &stdout
	cmd.Stdin = &stdin
	cmd.Stderr = &stderr

	host := fmt.Sprintf("1password-%s", account)

	stdin.Write([]byte(fmt.Sprintf("host=%s\n", host)))
	stdin.Write(eof)

	if err := cmd.Run(); err != nil {
		return "", errors.New(stderr.String()) // nolint:goerr113 // TODO: refactor
	}

	data, err := ReadInput(bytes.NewReader(stdout.Bytes()))

	if err != nil {
		return "", err
	}

	return data["password"], nil
}

// StoreInCache stores a 1password session token inside git cache.
func StoreInCache(account, token string, timeout uint) error {
	var stdout bytes.Buffer

	var stdin bytes.Buffer

	var stderr bytes.Buffer

	t := fmt.Sprintf("%d", timeout)
	cmd := exec.Command("git", "credential-cache", "--timeout", t, "store")
	cmd.Stdout = &stdout
	cmd.Stdin = &stdin
	cmd.Stderr = &stderr

	host := fmt.Sprintf("1password-%s", account)

	stdin.Write([]byte(fmt.Sprintf("host=%s\n", host)))
	stdin.Write([]byte("username=session-token\n"))
	stdin.Write([]byte(fmt.Sprintf("password=%s\n", token)))
	stdin.Write(eof)

	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String()) // nolint:goerr113 // TODO: refactor
	}

	return nil
}
