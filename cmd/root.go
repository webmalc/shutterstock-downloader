package cmd

import (
	"github.com/spf13/cobra"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	rootCmd        *cobra.Command
	config         *Config
	downloadRunner Runner
}

// configShow show the configuration.
func (s *CommandRouter) download(_ *cobra.Command, _ []string) {
	s.downloadRunner.Run()
}

// Run the router.
func (s *CommandRouter) Run() {
	s.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "download",
			Short: "Download all your collection images",
			Run:   s.download,
		},
	)
	err := s.rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(
	downloadRunner Runner,
) CommandRouter {
	config := NewConfig()

	return CommandRouter{
		config:         config,
		downloadRunner: downloadRunner,
		rootCmd:        &cobra.Command{Use: "shutterstock-downloader"},
	}
}
