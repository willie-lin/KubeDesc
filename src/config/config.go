package Config

import (
	"os"
	"path/filepath"
)

//ConfigStrings
//func ConfigStrings() string{
//
//	return GetConfig()
//
//}

func GetConfig() string {
	str, _ := os.Getwd()
	kubeconfig := filepath.Join(str, "src/config")
	return kubeconfig
}