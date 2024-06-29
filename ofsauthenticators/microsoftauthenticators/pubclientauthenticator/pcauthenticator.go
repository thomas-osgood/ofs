package pubclientauthenticator

import "google.golang.org/grpc/metadata"

// function designed to utilize MSAL to contact Microsoft and
// verify the login/authorization request.
func (pca *PublicClientAuthenticator) ValidateUser(md metadata.MD) (err error) {
	return nil
}
