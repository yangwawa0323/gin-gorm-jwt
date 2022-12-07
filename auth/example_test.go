package auth_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

var conf = &oauth2.Config{
	ClientID:     "2b0fd37061dc061aaf4b",
	ClientSecret: "38e1a4bc79ab2adfcdadac2d1330cf0c4b00a410",
	Scopes:       []string{"user", "repo"},
	Endpoint: oauth2.Endpoint{
		TokenURL: "https://github.com/login/oauth/access_token",
		AuthURL:  "https://github.com/login/oauth/authorize",
	},
}

func Exampleconfig() {
	ctx := context.Background()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	client.Get("...")
}

func Test_ExampleConfig_customHTTP(t *testing.T) {
	ctx := context.Background()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	// use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	_ = client

}
