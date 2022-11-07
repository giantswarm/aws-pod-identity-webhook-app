package cmd

import (
	"io"
	"os"

	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/aws-pod-identity-webhook/pkg/project"
)

type Config struct {
	Logger micrologger.Logger

	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:           project.Name(),
		Short:         project.Description(),
		Long:          project.Description(),
		RunE:          r.Run,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	f.Init(c)

	return c, nil
}
