// module designed to define a Microsoft Public Client Authenticator that
// can be used by OFS to verify a user.
package pubclientauthenticator

import "google.golang.org/grpc/metadata"

// function designed to utilize MSAL to contact Microsoft and
// verify the login/authorization request.
//
// ref:
//
// https://learn.microsoft.com/en-us/entra/msal/go/
func (pca *PublicClientAuthenticator) ValidateUser(md metadata.MD) (err error) {
	return nil
}
