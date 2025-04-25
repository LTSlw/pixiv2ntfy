package ntfy

import "encoding/base64"

type AuthUserPassword struct {
	Username string
	Password string
}

func (a *AuthUserPassword) AuthHeader() string {
	up := a.Username + ":" + a.Password
	return "Base " + base64.StdEncoding.EncodeToString([]byte(up))
}

type AuthToken struct {
	Token string
}

func (a *AuthToken) AuthHeader() string {
	return "Bearer " + a.Token
}

type Auth interface {
	AuthHeader() string
}
