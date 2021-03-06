package action

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// Git runs git commands inside the store or mounts
func (s *Action) Git(c *cli.Context) error {
	store := c.String("store")
	return s.Store.Git(store, c.Args()...)
}

// GitInit initializes a git repo including basic configuration
func (s *Action) GitInit(c *cli.Context) error {
	store := c.String("store")
	sk := c.String("sign-key")

	return s.gitInit(store, sk)
}

func (s *Action) gitInit(store, sk string) error {
	if sk == "" {
		s, err := s.askForPrivateKey(color.CyanString("Please select a key for signing Git Commits"))
		if err == nil {
			sk = s
		}
	}

	// for convenience, set defaults to user-selected values from available private keys
	// NB: discarding returned error since this is merely a best-effort look-up for convenience
	userName, userEmail, _ := s.askForGitConfigUser()

	userName, err := s.askForString(color.CyanString("Please enter a user name for password store git config"), userName)
	if err != nil {
		return err
	}
	userEmail, err = s.askForString(color.CyanString("Please enter an email address for password store git config"), userEmail)
	if err != nil {
		return err
	}

	if err := s.Store.GitInit(store, sk, userName, userEmail); err != nil {
		return err
	}
	fmt.Fprintln(color.Output, color.GreenString("Git initialized"))
	return nil
}
