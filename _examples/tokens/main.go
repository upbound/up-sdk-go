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

	"github.com/google/uuid"
	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/tokens"
)

// Creates a control plane token for the specified control plane.
func main() {
	// Prompt for auth input
	fmt.Println("Enter username: ")
	var user string
	fmt.Scanln(&user)
	fmt.Println("Enter password: ")
	var password string
	fmt.Scanln(&password)
	fmt.Println("Enter control plane ID: ")
	var cpID string
	fmt.Scanln(&cpID)

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
	client := tokens.NewClient(cfg)
	fmt.Println("Creating control plane token...")
	cp, err := client.Create(context.Background(), &tokens.TokenCreateParameters{
		Attributes: tokens.TokenAttributes{
			Name: "cool token",
		},
		Relationships: tokens.TokenRelationships{
			Owner: tokens.TokenOwner{
				Data: tokens.TokenOwnerData{
					Type: tokens.TokenOwnerControlPlane,
					ID:   uuid.MustParse(cpID),
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %v\n", cp.DataSet.Meta["jwt"])
}
