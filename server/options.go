package server

import "github.com/peaklyio/api-server/db/mongo"

type ServerOptions struct {
	BindAddress     string
	BindPort        int
	Domain          string
	EnableTLS       bool
	TLSOptions      *TLSOptions
	DatabaseOptions *DatabaseOptions
}

type TLSOptions struct {
}

type DatabaseOptions struct {
	MongoOptions *mongo.MongoOptions
}
