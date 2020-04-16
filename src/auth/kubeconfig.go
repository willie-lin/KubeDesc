package auth

import (
	authApi "KubeDesc/src/auth/api"
	"KubeDesc/src/errors"

	yaml "gopkg.in/yaml.v2"
	"k8s.io/client-go/tools/clientcmd/api"
)

type contextInfo struct {
	User string `yaml:"user"`
}

type contextEntry struct {
	Name    string      `yaml:"name"`
	Context contextInfo `yaml:"context"`
}

type userEntry struct {
	Name string   `yaml:"name"`
	User userInfo `yaml:"user"`
}

type authProviderConfig struct {
	AccessToken string `yaml:"access-token"`
}

type authProviderInfo struct {
	Config authProviderConfig `yaml:"config"`
}

type userInfo struct {
	AuthProvider authProviderInfo `yaml:"auth-provider"`
	Token        string           `yaml:"token"`
	Username     string           `yaml:"username"`
	Password     string           `yaml:"password"`
}

type kubeConfig struct {
	Contexts       []contextEntry `yaml:"contexts"`
	CurrentContext string         `yaml:"current-context"`
	Users          []userEntry    `yaml:"users"`
}
type kubeConfigAuthenticator struct {
	fileContent []byte
	authModes   authApi.AuthenticationModes
}

func (self *kubeConfigAuthenticator) GetAuthInfo() (api.AuthInfo, error) {
	kubeConfig, err := self.parseKubeConfig(self.fileContent)
	if err != nil {
		return api.AuthInfo{}, err
	}

	info, err := self.getCurrentUserInfo(*kubeConfig)
	if err != nil {
		return api.AuthInfo{}, err
	}
	return self.getAuthInfo(info)

}

func (self *kubeConfigAuthenticator) parseKubeConfig(bytes []byte) (*kubeConfig, error) {
	kubeConfig := new(kubeConfig)
	if err := yaml.Unmarshal(bytes, kubeConfig); err != nil {
		return nil, err
	}
	return kubeConfig, nil
}

func (self *kubeConfigAuthenticator) getCurrentUserInfo(config kubeConfig) (userInfo, error) {
	userName := ""
	for _, context := range config.Contexts {
		if context.Name == config.CurrentContext {
			userName = context.Context.User
		}
	}
	if len(userName) == 0 {
		return userInfo{}, errors.NewInvalid("Context matching current context not found. Check if your config file is valid.")
	}

	for _, user := range config.Users {
		if user.Name == userName {
			return user.User, nil
		}
	}
	return userInfo{}, errors.NewInvalid("User matching current context user not found. Check if your config file is valid.")
}

func (self *kubeConfigAuthenticator) getAuthInfo(info userInfo) (api.AuthInfo, error) {
	if len(info.Token) == 0 {
		info.Token = info.AuthProvider.Config.AccessToken
	}

	if len(info.Token) == 0 && (len(info.Password) == 0 || len(info.Username) == 0) {
		return api.AuthInfo{}, errors.NewInvalid("Not enough data to create auth info structure.")
	}
	result := api.AuthInfo{}
	if self.authModes.IsEnabled(authApi.Token) {
		result.Token = info.Token
	}

	if self.authModes.IsEnabled(authApi.Basic) {
		result.Username = info.Username
		result.Password = info.Password
	}
	return result, nil
}

func NewKubeConfigAuthenticator(spec *authApi.LoginSpec, authModes authApi.AuthenticationModes) authApi.Authenticator {
	return &kubeConfigAuthenticator{
		fileContent: []byte(spec.KubeConfig),
		authModes:   authModes,
	}
}
