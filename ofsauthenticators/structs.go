package ofsauthenticators

// struct used to pass a token credential.
type TokenRequest struct {
	Token string `json:"token" xml:"token"`
}

// struct used to pass a username-password credential.
type UserPassRequest struct {
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

// struct used to define the metadata headers used
// by the authenticators to get the username, password, etc.
type MetadataInfo struct {
	// header holding the username data.
	HDRUsername string `json:"userheader" xml:"userheader"`
	// header holding the password (or token) data.
	HDRPassword string `json:"passheader" xml:"passheader"`
}
