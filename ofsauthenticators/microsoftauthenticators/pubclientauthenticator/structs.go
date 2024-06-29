package pubclientauthenticator

// reference:
//
// https://learn.microsoft.com/en-us/entra/msal/go/
type PublicClientAuthenticator struct {
	authority    string
	clientid     string
	clientsecret string
	scope        []string
}

type PubClientAuthOption struct {
	// URL that will be used to verify the credentials.
	Authority string
	// ClientID that will be used when creating the public client.
	Clientid string
	// Secret that will be used when creating the public client.
	Clientsecret string
	// Permissions for the public client.
	Scope []string
}
