package user

import (
	"fmt"
	"os"
	"testing"

	peaklyclient "github.com/peaklyio/api-server/client"
	"github.com/peaklyio/api-server/e2e"
)

var (
	details *e2e.E2EEnvironmentDetails
	client
)

func TestMain(m *testing.M) {
	env := e2e.NewE2EEnvironment()
	d, err := env.SetUp()
	if err != nil {
		fmt.Printf("Major failure setting up environment: %v\n", err)
		os.Exit(99)
	}
	defer env.TearDown()
	details = d
	simpleAuth := &peaklyclient.SimpleAuthorization{}
	cfg := &peaklyclient.ServerConfiguration{
		Address: "localhost",
		Port:    details.ServerPort,
	}
	client, err := peaklyclient.NewClient(simpleAuth, cfg)
	if err != nil {
		t.Errorf("error loading new client: %v", err)
	}
	status := m.Run()
	fmt.Printf("Testing exit code [%d]\n", status)

	os.Exit(status)
}

func TestNewUser(t *testing.T) {


}
