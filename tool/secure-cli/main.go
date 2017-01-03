package main

import (
	"fmt"
	goaclient "github.com/goadesign/goa/client"
	"github.com/jamesallain/goa-fhir/client"
	"github.com/jamesallain/goa-fhir/tool/cli"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "Secure-cli",
		Short: `CLI client for the Secure service (http://goa.design/getting-started.html)`,
	}

	// Create client struct
	httpClient := newHTTPClient()
	c := client.New(goaclient.HTTPClientDoer(httpClient))

	// Register global flags
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "localhost:8088", "API hostname")
	app.PersistentFlags().DurationVarP(&httpClient.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")

	// Register signer flags
	var user, pass string
	app.PersistentFlags().StringVar(&user, "user", "", "Username used for authentication")
	app.PersistentFlags().StringVar(&pass, "pass", "", "Password used for authentication")
	var key, format string
	app.PersistentFlags().StringVar(&key, "key", "", "API key used for authentication")
	app.PersistentFlags().StringVar(&format, "format", "Bearer %s", "Format used to create auth header or query from key")

	// Parse flags and setup signers
	app.ParseFlags(os.Args)
	basicAuthSigner := newBasicAuthSigner(user, pass)
	jwtSigner := newJWTSigner(key, format)
	signinBasicAuthSigner := newSigninBasicAuthSigner(user, pass)

	// Initialize API client
	c.SetBasicAuthSigner(basicAuthSigner)
	c.SetJWTSigner(jwtSigner)
	c.SetSigninBasicAuthSigner(signinBasicAuthSigner)
	c.UserAgent = "Secure-cli/0"

	// Register API commands
	cli.RegisterCommands(app, c)

	// Execute!
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}

// newHTTPClient returns the HTTP client used by the API client to make requests to the service.
func newHTTPClient() *http.Client {
	// TBD: Change as needed (e.g. to use a different transport to control redirection policy or
	// disable cert validation or...)
	return http.DefaultClient
}

// newBasicAuthSigner returns the request signer used for authenticating
// against the basic_auth security scheme.
func newBasicAuthSigner(user, pass string) goaclient.Signer {
	return &goaclient.BasicSigner{
		Username: user,
		Password: pass,
	}

}

// newJWTSigner returns the request signer used for authenticating
// against the jwt security scheme.
func newJWTSigner(key, format string) goaclient.Signer {
	return &goaclient.APIKeySigner{
		SignQuery: false,
		KeyName:   "Authorization",
		KeyValue:  key,
		Format:    format,
	}

}

// newSigninBasicAuthSigner returns the request signer used for authenticating
// against the SigninBasicAuth security scheme.
func newSigninBasicAuthSigner(user, pass string) goaclient.Signer {
	return &goaclient.BasicSigner{
		Username: user,
		Password: pass,
	}

}
