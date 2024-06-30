package pubclientauthenticator

import (
	"fmt"
	"strings"

	maconsts "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/constants"
	madefaults "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/defaults"
	mamessages "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/messages"
)

// function designed to create, initialize and return a new
// public client authenticator object.
func NewPublicClientAuthenticator(opts ...PubClientAuthOptFunc) (authenticator *PublicClientAuthenticator, err error) {
	var curopt PubClientAuthOptFunc
	var defaults PubClientAuthOption = PubClientAuthOption{
		HeaderInfo: AuthHeaders{
			HdrToken:    madefaults.DEFAULT_HEADER_TOKEN,
			HdrUsername: madefaults.DEFAULT_HEADER_USERNAME,
		},
	}

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

	authenticator.authUrl = fmt.Sprintf(maconsts.AUTHORITY_FORMAT, defaults.Tenantid, defaults.Clientid)
	authenticator.headerInfo = defaults.HeaderInfo
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
	var curscope string
	var tmpscope []string = make([]string, 0)

	for _, curscope = range scope {
		curscope = strings.TrimSpace(curscope)
		if len(curscope) < 1 {
			continue
		}
		tmpscope = append(tmpscope, curscope)
	}

	scope = tmpscope

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
