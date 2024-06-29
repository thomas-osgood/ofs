package pubclientauthenticator

import (
	"fmt"
	"strings"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	maconsts "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/constants"
	mamessages "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/messages"
)

// function designed to create, initialize and return a new
// public client authenticator object.
func NewPublicClientAuthenticator(opts ...PubClientAuthOptFunc) (authenticator *PublicClientAuthenticator, err error) {
	var authority string
	var curopt PubClientAuthOptFunc
	var defaults PubClientAuthOption = PubClientAuthOption{}

	for _, curopt = range opts {
		err = curopt(&defaults)
		if err != nil {
			return nil, err
		}
	}

	err = defaults.validate()
	if err != nil {
		return nil, err
	}

	authenticator = new(PublicClientAuthenticator)

	// construct the authority url based on the tenant passed in.
	authority = fmt.Sprintf(maconsts.AUTHORITY_FORMAT, defaults.Tenantid)

	// create the public client application to use for authentication.
	authenticator.app, err = public.New(defaults.Clientid, public.WithAuthority(authority))
	if err != nil {
		return nil, err
	}

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

// set the authorization scope to for the user.
func WithScope(scope []string) PubClientAuthOptFunc {
	return func(pcao *PubClientAuthOption) error {
		if len(scope) < 1 {
			return fmt.Errorf(mamessages.ERR_SCOPE_NULL)
		}
		pcao.Scope = scope
		return nil
	}
}

// set the tenant to use when contacting Microsoft.
func WithTenantID(tenantid string) PubClientAuthOptFunc {
	return func(pcao *PubClientAuthOption) error {
		tenantid = strings.TrimSpace(tenantid)
		if len(tenantid) < 1 {
			return fmt.Errorf(mamessages.ERR_TENANTID_NULL)
		}
		pcao.Tenantid = tenantid
		return nil
	}
}
