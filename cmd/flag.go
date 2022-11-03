package cmd

import (
	goflag "flag"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type flag struct {
}

func (f *flag) Init(cmd *cobra.Command) {
	// Add command line flags for glog.
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

func (f *flag) Validate(cmd *cobra.Command) error {
	return nil
}
