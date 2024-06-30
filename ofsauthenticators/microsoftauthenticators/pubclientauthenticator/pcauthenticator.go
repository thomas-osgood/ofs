// module designed to define a Microsoft Public Client Authenticator that
// can be used by OFS to verify a user.
package pubclientauthenticator

import (
	"google.golang.org/grpc/metadata"
)

// function designed to utilize MSAL to contact Microsoft and
// verify the login/authorization request.
//
// ref:
//
// https://learn.microsoft.com/en-us/entra/msal/go/
//
// https://learn.microsoft.com/en-us/answers/questions/793793/azure-ad-validate-access-token
//
// https://www.voitanos.io/blog/validating-entra-id-generated-oauth-tokens/
func (pca *PublicClientAuthenticator) ValidateUser(md metadata.MD) (err error) {
	var pubkey string
	var token string

	token, err = pca.readMetadata(md)
	if err != nil {
		return err
	}

	pubkey, err = pca.readPublicKey()
	if err != nil {
		return err
	}

	err = pca.validJWT(token, pubkey)
	if err != nil {
		return err
	}

	return nil
}
