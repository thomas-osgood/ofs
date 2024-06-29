package pubclientauthenticator

import (
	"fmt"
	"strings"

	mamessages "github.com/thomas-osgood/ofs/ofsauthenticators/microsoftauthenticators/internal/messages"
)

func (pcao *PubClientAuthOption) validate() (err error) {
	var curscope string
	var tmpscope []string

	pcao.Clientid = strings.TrimSpace(pcao.Clientid)
	pcao.Tenantid = strings.TrimSpace(pcao.Tenantid)

	if len(pcao.Clientid) < 1 {
		return fmt.Errorf(mamessages.ERR_CLIENTID_NULL)
	}

	if len(pcao.Tenantid) < 1 {
		return fmt.Errorf(mamessages.ERR_TENANTID_NULL)
	}

	for _, curscope = range pcao.Scope {
		curscope = strings.TrimSpace(curscope)
		if len(curscope) < 1 {
			continue
		}
		tmpscope = append(tmpscope, curscope)
	}

	if len(tmpscope) < 1 {
		return fmt.Errorf(mamessages.ERR_SCOPE_NULL)
	}

	pcao.Scope = tmpscope

	return nil
}
