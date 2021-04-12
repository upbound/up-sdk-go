package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/controlplanes"
)

// Creates a control plane in the authenticated user's namespace, then gets the
// control plane using the returned ID.
func main() {
	// Prompt for auth input
	fmt.Println("Enter username: ")
	var user string
	fmt.Scanln(&user)
	fmt.Println("Enter password: ")
	var password string
	fmt.Scanln(&password)

	// Temporarily use dev API
	base, _ := url.Parse("https://api.dev-deba7a0e.u6d.dev")
	cj, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	httpClient := &up.HTTPClient{
		BaseURL: base,
		HTTP: &http.Client{
			Jar: cj,
		},
		UserAgent:    "up-sdk-go",
		ErrorHandler: &up.DefaultErrorHandler{},
	}
	cfg := up.NewConfig(func(cfg *up.Config) {
		cfg.Client = httpClient
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
	req, err := http.NewRequest("POST", "https://api.dev-deba7a0e.u6d.dev/v1/login", bytes.NewReader(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	err = cfg.Client.Do(req, nil)
	if err != nil {
		panic(err)
	}
	client := controlplanes.NewControlPlanesClient(cfg)
	fmt.Println("Creating control plane...")
	cp, err := client.Create(context.Background(), &controlplanes.ControlPlaneCreateParameters{
		Name:        "test",
		Namespace:   user,
		Description: "An example control plane.",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Permission: %s\n", cp.ControlPlanePermission)
	fmt.Printf("Status: %s\n", cp.ControlPlaneStatus)
	fmt.Printf("Info: %v\n", cp.ControlPlane)
	fmt.Println("Getting control plane...")
	res, err := client.Get(context.Background(), cp.ControlPlane.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Info: %v\n", res.ControlPlane)
}
