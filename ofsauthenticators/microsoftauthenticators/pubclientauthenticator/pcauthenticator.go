// module designed to define a Microsoft Public Client Authenticator that
// can be used by OFS to verify a user.
package pubclientauthenticator

import (
	"context"

	madefaults "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/defaults"
	"google.golang.org/grpc/metadata"
)

// function designed to utilize MSAL to contact Microsoft and
// verify the login/authorization request.
//
// ref:
//
// https://learn.microsoft.com/en-us/entra/msal/go/
func (pca *PublicClientAuthenticator) ValidateUser(md metadata.MD) (err error) {
	var cancel context.CancelFunc
	var ctx context.Context
	var password string
	var username string

	username, password, err = pca.readMetadata(md)
	if err != nil {
		return err
	}

	ctx, cancel = context.WithTimeout(context.Background(), madefaults.DEFAULT_TIMEOUT)
	defer cancel()

	_, err = pca.app.AcquireTokenByUsernamePassword(ctx, pca.scope, username, password)
	if err != nil {
		return err
	}

	return nil
}
