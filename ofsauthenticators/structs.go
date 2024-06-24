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
