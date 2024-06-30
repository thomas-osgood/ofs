package pubclientauthenticator

// reference:
//
// https://learn.microsoft.com/en-us/entra/msal/go/
type PublicClientAuthenticator struct {
	authUrl    string
	headerInfo AuthHeaders
	scope      []string
}

type PubClientAuthOption struct {
	// ClientID that will be used when creating the public client.
	Clientid string
	// Headers to read from metadata.
	HeaderInfo AuthHeaders
	// Secret that will be used when creating the public client.
	Tenantid string
	// Permissions for the public client.
	Scope []string
}

// struct defining the header names for the username and
// password headers that will be read from the metadata.
type AuthHeaders struct {
	HdrToken    string
	HdrUsername string
}
