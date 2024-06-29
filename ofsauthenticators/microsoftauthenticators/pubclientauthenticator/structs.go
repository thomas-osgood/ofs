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

// struct defining the header names for the username and
// password headers that will be read from the metadata.
//
// note: "password" could mean a password/passphrase or
// a token to use for authentication.
type AuthHeaders struct {
	HdrPassword string
	HdrUsername string
}
