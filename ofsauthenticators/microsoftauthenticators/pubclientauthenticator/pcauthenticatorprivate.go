package pubclientauthenticator

import (
	"encoding/base64"
	"fmt"
	"log"
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

	// TODO: parse response body and extract returned public key information.
	// this public key info will, most likely, be Base64 encoded.

	return pubkey, nil
}

// function designed to determine whether a given JWT is valid based on the
// JWT and public key provided.
//
// reference:
//
// https://stackoverflow.com/questions/51834234/i-have-a-public-key-and-a-jwt-how-do-i-check-if-its-valid-in-go
func (pca *PublicClientAuthenticator) validJWT(jwt string, pubkey string) (err error) {
	var jwtbody []byte
	var jwtheader []byte
	var jwtsig []byte
	var jwtsplit []string = strings.Split(jwt, ".")

	if len(jwtsplit) != 3 {
		return fmt.Errorf(mamessages.ERR_JWT_INVALID)
	}

	jwtheader, err = base64.RawURLEncoding.DecodeString(jwtsplit[0])
	if err != nil {
		return err
	}

	jwtbody, err = base64.RawURLEncoding.DecodeString(jwtsplit[1])
	if err != nil {
		return err
	}

	jwtsig, err = base64.RawURLEncoding.DecodeString(jwtsplit[2])
	if err != nil {
		return err
	}

	log.Printf("Header: %s\n", string(jwtheader))
	log.Printf("Body: %s\n", string(jwtbody))
	log.Printf("Sig Len: %d\n", len(jwtsig))

	return nil
}
