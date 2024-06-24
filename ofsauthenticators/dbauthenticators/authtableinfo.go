package dbauthenticators

// function designed to determine if the AuthTableInfo
// struct is nil. this will be set to true if all fields
// are empty strings.
func (ati *AuthTableInfo) IsNil() bool {
	var curcol string
	var fields []string

	ati.cleanStrings()

	fields = []string{ati.Tablename, ati.Usercolumn, ati.Passcolumn}

	for _, curcol = range fields {
		if len(curcol) > 0 {
			return false
		}
	}

	return true
}
