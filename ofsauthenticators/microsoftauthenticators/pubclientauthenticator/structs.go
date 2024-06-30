package pubclientauthenticator

import "time"

// reference:
//
// https://learn.microsoft.com/en-us/entra/msal/go/
type PublicClientAuthenticator struct {
	authUrl    string
	headerInfo AuthHeaders
	reqTimeout time.Duration
	scope      []string
}

type PubClientAuthOption struct {
	// ClientID that will be used when creating the public client.
	Clientid string
	// Headers to read from metadata.
	HeaderInfo AuthHeaders
	// Timeout for HTTP requests.
	ReqTimeout time.Duration
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

// struct defining the values that are of interest in the JWT header.
type JWTHeader struct {
	Alg string `json:"alg" xml:"alg"`
	Kid string `json:"kid" xml:"kid"`
	Typ string `json:"typ" xml:"typ"`
	X5t string `json:"x5t" xml:"x5t"`
}
