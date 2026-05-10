package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const userEnv = "LIBRE_LINKUP_USER"
const userPass = "LIBRE_LINKUP_PASS"
const libreProduct = "llu.android"
const libreVersion = "4.16.0"
const baseURL = "https://api-us.libreview.io"

func main() {
	slog.Info("Starting sugarctl")

	if err := godotenv.Load(); err != nil {
		slog.Error("Failed to load environment", "err", err)
		return
	}
	slog.Info("Environment loaded")

	userInput, passInput := os.Getenv(userEnv), os.Getenv(userPass)
	if strings.TrimSpace(userInput) == "" || strings.TrimSpace(passInput) == "" {
		slog.Error("User or password not provided")
		return
	}
	slog.Info("User and password gathered",
		"user_len", len(userInput),
		"pass_len", len(passInput))

	acc, err := login(userInput, passInput)
	if err != nil {
		slog.Error("Login failed", "error", err)
		return
	}
	slog.Info("Login successful", "account", acc)

}

type account struct {
	accountID string
	token     string
}

// login takes the user and pass and returns the token and account id
func login(user, pass string) (account, error) {
	acc := account{}
	u, err := url.Parse(baseURL + "/llu/auth/login")
	if err != nil {
		return acc, fmt.Errorf("Failed to parse login URL: %w", err)
	}
	bodyMap := map[string]string{"email": user, "password": pass}
	b, err := json.Marshal(bodyMap)
	if err != nil {
		return acc, fmt.Errorf("Failed to marshal body: %w", err)
	}
	r, err := http.NewRequest(
		http.MethodPost,
		u.String(),
		bytes.NewBuffer(b))
	//r.Header.Set("accept-encoding", "gzip")
	r.Header.Set("cache-control", "no-cache")
	r.Header.Set("connection", "Keep-Alive")
	r.Header.Set("product", libreProduct)
	r.Header.Set("version", libreVersion)
	r.Header.Set("Content-Type", "application/json")

	// make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return acc, fmt.Errorf("Failed to make request: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return acc, fmt.Errorf("Failed to read response body: %w", err)
	}

	lr := loginResponse{}
	if err := json.Unmarshal(body, &lr); err != nil {
		return acc, fmt.Errorf("Failed to unmarshal response: %w", err)
	}

	acc.token = lr.Data.AuthTicket.Token
	acc.accountID = lr.Data.User.ID
	return acc, nil
}

type connections struct{}

func getConnections(acc account) (connections, error) {
	cons := connections{}
	return cons, errors.New("Implementation not finished")
}
