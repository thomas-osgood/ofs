package queries

// query designed to select the username,password columns from
// the auth table where the username = the parameter. the username
// comparison in this query is CASE sensitive.
const SELECT_USER_PASSWORD string = `select %s,%s from %s where %s=$1`
