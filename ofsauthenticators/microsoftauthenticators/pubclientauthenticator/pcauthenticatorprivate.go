package pubclientauthenticator

import (
	"fmt"
	"strings"

	mamessages "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/messages"
	"google.golang.org/grpc/metadata"
)

// function designed to extract the username and password information
// from the passed in metadata.
func (pca *PublicClientAuthenticator) readMetadata(md metadata.MD) (username string, password string, err error) {

	password = strings.TrimSpace(md.Get(pca.headerInfo.HdrPassword)[0])
	if len(password) < 1 {
		return "", "", fmt.Errorf(mamessages.ERR_PASSWORD_NULL)
	}

	username = strings.TrimSpace(md.Get(pca.headerInfo.HdrUsername)[0])
	if len(username) < 1 {
		return "", "", fmt.Errorf(mamessages.ERR_USERNAME_NULL)
	}

	return username, password, nil
}
