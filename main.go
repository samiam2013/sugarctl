package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

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

	slog.Info("waiting 10 seconds for the api servers to sync")
	time.Sleep(10 * time.Second)

	cons, err := getConnections(acc)
	if err != nil {
		slog.Error("Failed getting connections: %w", err)
		return
	}
	_ = cons

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
	r.Header.Set("Cache-Control", "No-Cache")
	r.Header.Set("Connection", "Keep-Alive")
	r.Header.Set("Product", libreProduct)
	r.Header.Set("Version", libreVersion)
	r.Header.Set("Content-Type", "application/json")

	// make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return acc, fmt.Errorf("Failed to make request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
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
	u, err := url.Parse(baseURL + "/llu/connections")
	if err != nil {
		return cons, fmt.Errorf("Failed to parse URL for request: %w", err)
	}

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return cons, fmt.Errorf("Failed to create request for request: %w", err)
	}
	sha256id := sha256.New()
	sha256id.Write([]byte(acc.accountID))
	hexIDhash := hex.EncodeToString(sha256id.Sum(nil))
	r.Header.Set("Authorization", "Bearer "+acc.token)
	r.Header.Set("Account-ID", string(hexIDhash))
	r.Header.Set("Cache-Control", "No-Cache")
	r.Header.Set("Connection", "Keep-Alive")
	r.Header.Set("Product", libreProduct)
	r.Header.Set("Version", libreVersion)
	r.Header.Set("Content-Type", "application/json")

	c := http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		return cons, fmt.Errorf("Failed to do request for request: %w", err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return cons, fmt.Errorf("Failed to read response body for request: %w", err)
	}
	fmt.Println(string(b))

	return cons, errors.New("Implementation not finished")
}
