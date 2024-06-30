package pubclientauthenticator

import (
	"fmt"
	"net/http"
	"strings"

	mamessages "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/messages"
	"google.golang.org/grpc/metadata"
)

// function designed to extract the username and password information
// from the passed in metadata.
func (pca *PublicClientAuthenticator) readMetadata(md metadata.MD) (token string, err error) {

	token = strings.TrimSpace(md.Get(pca.headerInfo.HdrToken)[0])
	if len(token) < 1 {
		return "", fmt.Errorf(mamessages.ERR_TOKEN_NULL)
	}

	return token, nil
}

// function designed to reach out to Microsoft and pull down the public key
// generated for the client. this will help verify a JWT.
func (pca *PublicClientAuthenticator) readPublicKey() (pubkey string, err error) {
	var client *http.Client = http.DefaultClient
	var req *http.Request
	var resp *http.Response

	client.Timeout = pca.reqTimeout

	req, err = http.NewRequest(http.MethodGet, pca.authUrl, nil)
	if err != nil {
		return "", err
	}

	resp, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return pubkey, nil
}

// function designed to determine whether a given JWT is valid based on the
// JWT and public key provided.
func (pca *PublicClientAuthenticator) validJWT(jwt string, pubkey string) (err error) {
	return nil
}
