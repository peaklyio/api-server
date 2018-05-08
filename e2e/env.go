package e2e

import (
	"os"

	"fmt"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/container"
	cliflags "github.com/docker/cli/cli/flags"
	"github.com/kubicorn/kubicorn/pkg/namer"
	"github.com/peaklyio/api-server/mongo"
	"github.com/peaklyio/api-server/server"
	"github.com/phayes/freeport"
)

type E2EEnvironment struct {
	details *E2EEnvironmentDetails
}

type E2EEnvironmentDetails struct {
	ContainerName string
	MongoPort     int
	ServerPort    int
}

func NewE2EEnvironment() *E2EEnvironment {
	e := &E2EEnvironment{}
	return e
}

func (e *E2EEnvironment) SetUp() (*E2EEnvironmentDetails, error) {

	name := namer.RandomName()
	mongoPort, err := freeport.GetFreePort()
	if err != nil {
		return nil, fmt.Errorf("unable to get free port: %v", err)
	}
	serverPort, err := freeport.GetFreePort()
	if err != nil {
		return nil, fmt.Errorf("unable to get free port: %v", err)
	}

	// --- Mongo ---
	cli := command.NewDockerCli(os.Stdin, os.Stdout, os.Stderr, true)
	opts := &cliflags.ClientOptions{
		Common: &cliflags.CommonOptions{},
	}
	cli.Initialize(opts)
	cobra := container.NewRunCommand(cli)
	cobra.Flags().Set("name", "name")
	cobra.Flags().Set("rm", "1")
	cobra.Flags().Set("d", "1")
	cobra.Flags().Set("p", fmt.Sprintf("%d:%d", mongoPort, mongoPort))
	cobra.RunE(cobra, []string{"mongo:3.2"})

	// --- Server ---
	options := &server.ServerOptions{
		Domain:      "test.peakly",
		BindPort:    serverPort,
		BindAddress: "0.0.0.0",
		DatabaseOptions: &server.DatabaseOptions{
			MongoOptions: &mongo.MongoOptions{
				Database: "localhost",
				Port:     mongoPort,
			},
		},
	}
	go func() {
		err := server.ListenAndServe(options)
		if err != nil {
			fmt.Println(fmt.Errorf("error running server: %v", err))
		}
	}()

	details := &E2EEnvironmentDetails{
		ContainerName: name,
		MongoPort:     mongoPort,
		ServerPort:    serverPort,
	}
	e.details = details
	return details, nil
}

func (e *E2EEnvironment) TearDown() error {
	cli := command.NewDockerCli(os.Stdin, os.Stdout, os.Stderr, true)
	opts := &cliflags.ClientOptions{
		Common: &cliflags.CommonOptions{},
	}
	cli.Initialize(opts)
	cobra := container.NewRunCommand(cli)
	container.NewStopCommand(cli)
	cobra.RunE(cobra, []string{e.details.ContainerName})
	return nil
}
