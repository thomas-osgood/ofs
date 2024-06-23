package ofsauthenticators

// generic interface defining an authentication request.
type AuthRequest interface {
	TokenRequest | UserPassRequest
}
