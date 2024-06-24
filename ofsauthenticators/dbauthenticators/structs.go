package dbauthenticators

type AuthTableInfo struct {
	// name of the table that holds the verification info.
	Tablename string
	// column name of the column that holds the username.
	Usercolumn string
	// column name of the column that holds the password.
	Passcolumn string
}
