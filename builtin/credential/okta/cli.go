package okta

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
	pwd "github.com/hashicorp/vault/helper/password"
)

// CLIHandler struct
type CLIHandler struct{}

// Auth cli method
func (h *CLIHandler) Auth(c *api.Client, m map[string]string) (*api.Secret, error) {
	mount, ok := m["mount"]
	if !ok {
		mount = "okta"
	}

	username, ok := m["username"]
	if !ok {
		return nil, fmt.Errorf("'username' var must be set")
	}
	password, ok := m["password"]
	if !ok {
		fmt.Printf("Password (will be hidden): ")
		var err error
		password, err = pwd.Read(os.Stdin)
		fmt.Println()
		if err != nil {
			return nil, err
		}
	}

	data := map[string]interface{}{
		"password": password,
	}

	path := fmt.Sprintf("auth/%s/login/%s", mount, username)
	secret, err := c.Logical().Write(path, data)
	if err != nil {
		return nil, err
	}
	if secret == nil {
		return nil, fmt.Errorf("empty response from credential provider")
	}

	return secret, nil
}

// Help method for okta cli
func (h *CLIHandler) Help() string {
	help := `
Usage: vault login -method=okta [CONFIG K=V...]

  The OKTA authentication method allows users to authenticate using OKTA.

  Authenticate as "sally":

      $ vault login -method=okta username=sally
      Password (will be hidden):

  Authenticate as "bob":

      $ vault login -method=okta username=bob password=password

Configuration:

  password=<string>
      OKTA password to use for authentication. If not provided, the CLI will
      prompt for this on stdin.

  username=<string>
      OKTA username to use for authentication.
`

	return strings.TrimSpace(help)
}
