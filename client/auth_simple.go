package client

type SimpleAuthorization struct {
}

func (s *SimpleAuthorization) Authorize() error {
	return nil
}

func (s *SimpleAuthorization) Authenticate() error {
	return nil
}
