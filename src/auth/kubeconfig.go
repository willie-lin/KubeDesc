package auth

import (
	authApi "KubeDesc/src/auth/api"
	"gopkg.in/yaml.v2"

	_ "gopkg.in/yaml.v2"
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

func (self *kubeConfigAuthenticator) GetAuthInfo() (*api.AuthInfo, error) {
	kubeConfig, err := self.parseKubeConfig(self.fileContent)

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
		return userInfo{}, errors.NewInvalid
	}

	for _, user := range config.Users {

	}
}
