package microsoftauthenticators

import (
	"fmt"
	"strings"

	mamessages "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/messages"
)

// function designed to create, initialize and return a new
// public client authenticator object.
func NewPublicClientAuthenticator(opts ...PubClientAuthOptFunc) (authenticator *PublicClientAuthenticator, err error) {
	var curopt PubClientAuthOptFunc
	var defaults PubClientAuthOption = PubClientAuthOption{}

	for _, curopt = range opts {
		err = curopt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	authenticator = new(PublicClientAuthenticator)
	authenticator.authority = defaults.Authority
	authenticator.clientid = defaults.Clientid
	authenticator.clientsecret = defaults.Clientsecret
	authenticator.scope = defaults.Scope

	return authenticator, nil
}

// set the clientid to use when contacting Microsoft.
func WithClientID(clientid string) PubClientAuthOptFunc {
	return func(pcao *PubClientAuthOption) error {
		clientid = strings.TrimSpace(clientid)
		if len(clientid) < 1 {
			return fmt.Errorf(mamessages.ERR_CLIENTID_NULL)
		}
		pcao.Clientid = clientid
		return nil
	}
}
