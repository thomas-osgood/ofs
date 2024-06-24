package postgresauthenticator

import "strings"

// function designed to run the TrimSpace function on
// all fields within the AuthTableInfo struct.
func (ati *AuthTableInfo) cleanStrings() {
	ati.Tablename = strings.TrimSpace(ati.Tablename)
	ati.Tablename = strings.TrimSpace(ati.Passcolumn)
	ati.Tablename = strings.TrimSpace(ati.Usercolumn)
}
