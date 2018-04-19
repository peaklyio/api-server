package server

type ServerOptions struct {
	BindAddress     string
	BindPort        int
	EnableTLS       bool
	TLSOptions      *TLSOptions
	DatabaseOptions *DatabaseOptions
}

type TLSOptions struct {
}

type DatabaseOptions struct {
}
