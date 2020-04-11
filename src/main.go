package main

import (
	"github.com/spf13/pflag"
	"net"
)

var (
	argInsecurePort        = pflag.Int("insecure-port", 9090, "The port to listen to for incoming HTTP requests.")
	argPort                = pflag.Int("port", 8443, "The secure port to listen to for incoming HTTPS requests.")
	argInsecureBindAddress = pflag.IP("insecure-bind-address", net.IPv4(127, 0, 0, 1), "The IP address on which to serve the --insecure-port (set to 127.0.0.1 for all interfaces).")
	argBindAddress         = pflag.IP("bind-address", net.IPv4(0, 0, 0, 0), "The IP address on which to serve the --port (set to 0.0.0.0 for all interfaces).")
	argDefaultCertDir      = pflag.String("default-cert-dir", "/certs", "Directory path containing '--tls-cert-file' and '--tls-key-file' files. Used also when auto-generating certificates flag is set.")
	argCertFile            = pflag.String("tls-cert-file", "", "File containing the default x509 Certificate for HTTPS.")
	argKeyFile             = pflag.String("tls-key-file", "", "File containing the default x509 private key matching --tls-cert-file.")
	argApiserverHost       = pflag.String("apiserver-host", "", "The address of the Kubernetes Apiserver "+
		"to connect to in the format of protocol://address:port, e.g., "+
		"http://localhost:8080. If not specified, the assumption is that the binary runs inside a "+
		"Kubernetes cluster and local discovery is attempted.")
	argMetricsProvider = pflag.String("metrics-provider", "sidecar", "Select provider type for metrics. 'none' will not check metrics.")
	argHeapsterHost    = pflag.String("heapster-host", "", "The address of the Heapster Apiserver "+
		"to connect to in the format of protocol://address:port, e.g., "+
		"http://localhost:8082. If not specified, the assumption is that the binary runs inside a "+
		"Kubernetes cluster and service proxy will be used.")

	argKubeConfigFile = pflag.String("kubeconfig", "", "Path to kubeconfig file with authorization and master location information.")
)

func main() {

}
