package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/develerik/git-credential-1password/git"
	"github.com/develerik/git-credential-1password/onepassword"
)

// storeCmd represents the store operation.
var storeCmd = &cobra.Command{
	Use:              "store",
	TraverseChildren: true,
	Args:             cobra.NoArgs,
	Short:            "Store the credential",
	Run: func(cmd *cobra.Command, args []string) {
		if err := storeCredentials(os.Stdin); err != nil {
			if _, err = fmt.Fprintf(os.Stderr, "%s", err); err != nil {
				panic(err)
			}
			os.Exit(1)
		}
	},
}

func storeCredentials(r io.Reader) error {
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

	// only store complete credentials
	host, hasHost := data["host"]
	password, hasPassword := data["password"]
	protocol, hasProtocol := data["protocol"]
	username, hasUsername := data["username"]

	if !hasHost || !hasPassword || !hasProtocol || !hasUsername {
		return nil
	}

	return c.StoreCredentials(protocol, host, username, password)
}
