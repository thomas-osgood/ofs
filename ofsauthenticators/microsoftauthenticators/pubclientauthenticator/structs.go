package pubclientauthenticator

import "github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"

// reference:
//
// https://learn.microsoft.com/en-us/entra/msal/go/
type PublicClientAuthenticator struct {
	app   public.Client
	scope []string
}

type PubClientAuthOption struct {
	// ClientID that will be used when creating the public client.
	Clientid string
	// Secret that will be used when creating the public client.
	Tenantid string
	// Permissions for the public client.
	Scope []string
}
