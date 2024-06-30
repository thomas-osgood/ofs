// module designed to define a Microsoft Public Client Authenticator that
// can be used by OFS to verify a user.
package pubclientauthenticator

import (
	"log"

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

	log.Printf("Token: %s\n", token)
	log.Printf("Pubkey: %s\n", pubkey)

	return nil
}
