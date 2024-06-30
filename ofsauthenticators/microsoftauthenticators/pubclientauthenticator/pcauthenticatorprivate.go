package pubclientauthenticator

import (
	"fmt"
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
