package auth

import (
	authApi "KubeDesc/src/auth/api"
	"k8s.io/client-go/tools/clientcmd/api"
)

//  implements authenticator interface
type basicAuthenticator struct {
	username string
	password string
}

func (self *basicAuthenticator) GetAuthInfo() (api.AuthInfo, error) {
	return api.AuthInfo{
		Username: self.username,
		Password: self.password,
	}, nil
}

func NewBasicAuthenticator(spec *authApi.LoginSpec) authApi.Authenticator {
	return &basicAuthenticator{
		username: spec.Username,
		password: spec.Password,
	}
}
