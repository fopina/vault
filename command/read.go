package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// ReadCommand is a Command that reads data from the Vault.
type ReadCommand struct {
	Meta
}

func (c *ReadCommand) Run(args []string) int {
	flags := c.Meta.FlagSet("read", FlagSetDefault)
	flags.Usage = func() { c.Ui.Error(c.Help()) }
	if err := flags.Parse(args); err != nil {
		return 1
	}

	args = flags.Args()
	if len(args) != 1 {
		c.Ui.Error("read expects one argument")
		flags.Usage()
		return 1
	}
	path := args[0]

	client, err := c.Client()
	if err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error initializing client: %s", err))
		return 2
	}

	secret, err := client.Logical().Read(path)
	if err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error reading %s: %s", path, err))
		return 1
	}

	b, err := json.Marshal(secret)
	if err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error reading %s: %s", path, err))
		return 1
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	c.Ui.Output(out.String())
	return 0
}

func (c *ReadCommand) Synopsis() string {
	return "Read data or secrets from Vault"
}

func (c *ReadCommand) Help() string {
	helpText := `
Usage: vault read [options] path

  Read data from Vault.

  Read reads data at the given path from Vault. This can be used to
  read secrets and configuration as well as generate dynamic values from
  materialized backends. Please reference the documentation for the
  backends in use to determine key structure.

General Options:

  -address=TODO           The address of the Vault server.

  -ca-cert=path           Path to a PEM encoded CA cert file to use to
                          verify the Vault server SSL certificate.

  -ca-path=path           Path to a directory of PEM encoded CA cert files
                          to verify the Vault server SSL certificate. If both
                          -ca-cert and -ca-path are specified, -ca-path is used.

  -insecure               Do not verify TLS certificate. This is highly
                          not recommended. This is especially not recommended
                          for unsealing a vault.

`
	return strings.TrimSpace(helpText)
}
