package auth

import "context"

func NewAuthentication(clientId,clientSecret string) *Authentication {
	return &Authentication{
		clientID: clientId,
		clientSecret: clientSecret,
	}
}

type Authentication struct {
	clientID     string
	clientSecret string
}

func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (
	map[string]string, error) {
		return map[string]string{
			ClientHeaderKey: a.clientID,
			ClientSecretKey: a.clientSecret,
		}, nil
		
}




func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
