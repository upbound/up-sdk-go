package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/namespaces"
)

// Lists all namespaces the authenticated user has access to.
func main() {
	// Prompt for auth input
	fmt.Println("Enter username: ")
	var user string
	fmt.Scanln(&user)
	fmt.Println("Enter password: ")
	var password string
	fmt.Scanln(&password)

	var api string
	if api = os.Getenv("UP_ENDPOINT"); api == "" {
		api = "https://api.upbound.io"
	}

	base, _ := url.Parse(api)
	cj, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	upClient := up.NewClient(func(c *up.HTTPClient) {
		c.BaseURL = base
		c.HTTP = &http.Client{
			Jar: cj,
		}
	})
	cfg := up.NewConfig(func(cfg *up.Config) {
		cfg.Client = upClient
	})
	auth := &struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}{
		ID:       user,
		Password: password,
	}
	jsonStr, err := json.Marshal(auth)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, path.Join(api, "v1", "login"), bytes.NewReader(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	err = cfg.Client.Do(req, nil)
	if err != nil {
		panic(err)
	}
	client := namespaces.NewClient(cfg)
	ns, err := client.List(context.Background())
	if err != nil {
		panic(err)
	}
	for _, n := range ns {
		fmt.Println(n.Namespace.Name)
	}
}
