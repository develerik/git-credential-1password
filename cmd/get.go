package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/develerik/git-credential-1password/git"
	"github.com/develerik/git-credential-1password/onepassword"
)

// getCmd represents the get operation.
var getCmd = &cobra.Command{
	Use:              "get",
	TraverseChildren: true,
	Args:             cobra.NoArgs,
	Short:            "Return a matching credential, if any exists",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getCredentials(os.Stdin, os.Stdout); err != nil {
			if _, err = fmt.Fprintf(os.Stderr, "%s", err); err != nil {
				panic(err)
			}
			os.Exit(1)
		}
	},
}

func getCredentials(r io.Reader, w io.Writer) error {
	data, err := git.ReadInput(r)

	if err != nil {
		return err
	}

	c := onepassword.Client{
		Account: account,
	}

	if err := c.Login(cache); err != nil {
		return err
	}

	host, exist := data["host"]

	if !exist {
		return fmt.Errorf("missing host to check credentials: %v", data) // nolint:goerr113 // TODO: refactor
	}

	creds, err := c.GetCredentials(host)

	if err != nil {
		return err
	}

	data["username"] = creds.Username
	data["password"] = creds.Password

	buf := new(bytes.Buffer)

	for k, v := range data {
		if _, err = fmt.Fprintf(buf, "%s=%s\n", k, v); err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(w, buf.String())

	return err
}
