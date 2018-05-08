package client

type Authorization interface {
	Authenticate() error
	Authorize() error
}
