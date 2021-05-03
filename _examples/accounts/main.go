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

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/accounts"
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
	u, err := base.Parse("/v1/login")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	err = cfg.Client.Do(req, nil)
	if err != nil {
		panic(err)
	}
	client := accounts.NewClient(cfg)
	accounts, err := client.List(context.Background())
	if err != nil {
		panic(err)
	}
	for _, a := range accounts {
		fmt.Println(a.Account.Name)
	}
}
