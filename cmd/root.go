// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/kubicorn/kubicorn/pkg/logger"
	"github.com/peaklyio/api-server/server"
	"github.com/spf13/cobra"
)

var (
	o = &server.ServerOptions{}
)

var RootCmd = &cobra.Command{
	Use:   "api-server",
	Short: "The peaklyio API server",
	Long:  `The peaklyio API server`,
	Run: func(cmd *cobra.Command, args []string) {

		err := server.ListenAndServe(o)
		if err != nil {
			logger.Critical("Fatal error: %v", err)
			os.Exit(99)
		}
		os.Exit(1)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Critical("High level fatal error: %v", err)
		os.Exit(109)
	}
}

func init() {
	// TODO @kris-nova TLS needs to be fleshed out more
	RootCmd.Flags().BoolVarP(&o.EnableTLS, "enable-tls", "t", false, "Enable TLS for the server (Requires TLS flags to be set)")
	RootCmd.Flags().StringVarP(&o.BindAddress, "bind-address", "a", "0.0.0.0", "The bind address for the server")
	RootCmd.Flags().IntVarP(&o.BindPort, "bind-address", "a", 80, "The bind port for the server")
}
