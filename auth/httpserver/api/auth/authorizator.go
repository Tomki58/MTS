package auth

type Authorizator interface {
	CheckCredentials(username, password string) bool
}
