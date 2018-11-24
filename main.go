package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var httpsEnvVar = "HEALTHCHECK_HTTPS"
var portEnvVar = "HEALTHCHECK_PORT"

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	port, _ := os.LookupEnv(httpsEnvVar)
	if len(port) > 0 {
		port = fmt.Sprintf(":%s", port)
	}
	httpsEnv, found := os.LookupEnv(httpsEnvVar)
	https := ""
	if found && len(httpsEnv) != 0 && strings.ToUpper(httpsEnv)[0] == 't' {
		https = "s"
	}
	resp, err := http.Get(fmt.Sprintf(
		"http%s://127.0.0.1%s/health", https, port,
	))
	if err != nil {
		fmt.Println("Error while connecting:", err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		os.Exit(1)
	}

	os.Exit(0)
}
