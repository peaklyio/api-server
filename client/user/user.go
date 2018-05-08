package user

import "github.com/peaklyio/api-server/client"

type UserResource struct {
	client *client.Client
}

func NewUserResource(client *client.Client) *UserResource {
	return &UserResource{
		client: client,
	}
}
