package main

import (
	"context"
	"log"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/aws-pod-identity-webhook/v2/cmd"
)

func main() {
	err := mainE(context.Background())
	if err != nil {
		log.Fatalf("Error: %s\n", err)
		os.Exit(2)
	}
}

func mainE(ctx context.Context) error {
	var err error

	var logger micrologger.Logger
	{
		c := micrologger.Config{}

		logger, err = micrologger.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var rootCommand *cobra.Command
	{
		c := cmd.Config{
			Logger: logger,
		}
		rootCommand, err = cmd.New(c)
		if err != nil {
			return err
		}
	}

	err = rootCommand.ExecuteContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
