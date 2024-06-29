package pubclientauthenticator

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

// set the authority to use when contacting Microsoft.
func WithAuthority(authority string) PubClientAuthOptFunc {
	return func(pcao *PubClientAuthOption) error {
		authority = strings.TrimSpace(authority)
		if len(authority) < 1 {
			return fmt.Errorf(mamessages.ERR_AUTHORITY_NULL)
		}
		pcao.Authority = authority
		return nil
	}
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

// set the client secret to use when contacting Microsoft.
func WithClientSecret(clientsecret string) PubClientAuthOptFunc {
	return func(pcao *PubClientAuthOption) error {
		clientsecret = strings.TrimSpace(clientsecret)
		if len(clientsecret) < 1 {
			return fmt.Errorf(mamessages.ERR_CLIENTID_NULL)
		}
		pcao.Clientsecret = clientsecret
		return nil
	}
}
